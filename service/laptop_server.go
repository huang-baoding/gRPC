package service

import (
	"bytes"
	"context"
	"errors"
	"grpctest/pb"
	"io"
	"log"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const maxImageSize = 1 << 20 //图像的最大长度为1兆字节			1<<10是1KB;1<<20是1MB

//定义一个结构体封装server的方法
type LaptopServer struct {
	laptopStore LaptopStore //一个接口，里面有存储和查找函数
	imageStore  ImageStore
	ratingStore RatingStore
}

//返回一个&laptop
func NewLaptopServer(laptopStore LaptopStore, imageStore ImageStore, ratingStore RatingStore) *LaptopServer {
	return &LaptopServer{laptopStore, imageStore, ratingStore}
}

//一元rpc//////////////////////////////////////////////////
func (server *LaptopServer) CreateLaptop( ///////////////////////
	ctx context.Context,
	req *pb.CreateLaptopRequest,
) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop() //得到来自客户端的req（CreateLaptopRequest）中定义的laptop字段的值
	log.Printf("receive a create-laptop with id:%s", laptop.Id)

	if len(laptop.Id) > 0 {
		//检查是否是一个有效的uuid
		_, err := uuid.Parse(laptop.Id) //解析uuid，如果err不为空表示解析出错
		if err != nil {
			//status是rpc的状态码子包，可以返回一些常见错误
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid uuid:%v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			//codes.Internal表示是服务器内部的错误
			return nil, status.Errorf(codes.Internal, "can not generate a new laptop uuid:%v", err)
		}
		laptop.Id = id.String()
	}

	//假设在这里做一些繁重的处理。
	time.Sleep(2 * time.Second) //用来测试客户端的超时检测功能。
	if ctx.Err() == context.Canceled {
		log.Print("Request is cancel")
		return nil, status.Error(codes.Canceled, "request is canceled")
	}
	//实现客户端退出后，服务器不会继续保存laptop

	if err := contextError(ctx); err != nil {
		return nil, err
	}

	//一般我们在这时候应该把laptop存储到数据库中，（留着之后实现）
	//现在我们将laptop保存到store（内存）中
	err := server.laptopStore.Save(laptop) //server结构体实现了Store接口，此接口中有函数Save（）
	if err != nil {
		//检查是否因为记录已存在而出错
		code := codes.Internal                //(codes.Internal是服务器错误)/////////////////////////////////
		if errors.Is(err, ErrAlreadyExists) { ///////////////////
			code = codes.AlreadyExists //(已经存在)
		}
		//如果有错误，我们将错误返回给客户端
		return nil, status.Errorf(code, "cannot ssave laptop to the store:%v", err)
	}

	log.Printf("save laptop with id:%v", laptop.Id)

	//使用laptopid创建一个新的响应对象id,然后将此对象返回给调用者（client）
	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return res, nil
}

func (server *LaptopServer) SearchLaptop(
	req *pb.SearchLaptopRequest,
	stream pb.LaptopService_SearchLaptopServer) error {
	//第一件事是从请求中获取过滤器。
	filter := req.GetFilter()
	log.Printf("receive a search-laptop with filter:%w", filter) //找到了，记录日志

	err := server.laptopStore.Search(
		stream.Context(), //在流中获取上下文，将其传递给search函数
		filter,           //传入过滤器
		func(laptop *pb.Laptop) error { //传入回调函数
			res := &pb.SearchLaptopResponse{Laptop: laptop} //用此电脑创建一个新的响应对象

			err := stream.Send(res) //调用此函数可以将其发回客户端
			if err != nil {
				return err
			}

			log.Printf("sent laptop with id:%s", laptop.GetId()) //已发送，记录日志
			return nil
		},
	)
	if err != nil {
		return status.Errorf(codes.Internal, "unexpected error:%w", err) //发送内部错误，返回状态码
	}
	return nil
}

//以客户端流的方式上传电脑图片
func (server *LaptopServer) UploadImage(stream pb.LaptopService_UploadImageServer) error {
	req, err := stream.Recv() //接收一个包含stream信息的请求
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot receive image info"))
	}

	laptopID := req.GetInfo().GetLaptopId()   //获取电脑id
	imageType := req.GetInfo().GetImageType() //获取图像类型
	log.Printf("receive an upload-image request for laptop %s with image type %s", laptopID, imageType)

	laptop, err := server.laptopStore.Find(laptopID) //确保该id存在
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot find laptop: %v", err))
	}
	if laptop == nil {
		return logError(status.Errorf(codes.InvalidArgument, "laptop id %s doesn't exist", laptopID))
	}

	imageData := bytes.Buffer{} //创建一个字节缓冲区来存储图像
	imageSize := 0              //记录图像大小

	//循环接收图像数据
	for {
		//检查上下文
		if err := contextError(stream.Context()); err != nil {
			return err
		}

		log.Print("waiting to receive more data")

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}

		chunk := req.GetChunkData() //从请求中获取数据块
		size := len(chunk)          //获取数据块的长度

		log.Printf("receive a chunk with size: %d", size)

		imageSize += size             //图像总长度
		if imageSize > maxImageSize { //提前设置图片的最大长度
			return logError(status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", imageSize, maxImageSize))
		}

		//假设缓慢写入
		// time.Sleep(time.Second)

		_, err = imageData.Write(chunk) //将接受到的数据块附加到图像数据中
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
		}
	}

	//将图片数据保存到store，并取回图像id
	imageID, err := server.imageStore.Save(laptopID, imageType, imageData)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot save image to the store: %v", err))
	}

	res := &pb.UploadImageResponse{ //创建带有图像ID和图像大小的响应对象。
		Id:   imageID,
		Size: uint32(imageSize),
	}

	err = stream.SendAndClose(res) //将响应发送回到客户端。
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot send response: %v", err))
	}

	log.Printf("saved image with id: %s, size: %d", imageID, imageSize)
	return nil
}

//RateLaptop是一个双向流RPC，它允许客户端对笔记本电脑流进行评分，并返回每个笔记本电脑的平均分数流
func (server *LaptopServer) RateLaptop(stream pb.LaptopService_RateLaptopServer) error {
	for {	//因为要在流中接收多个请求，所以我们使用for循环
		//先检查上下文是否已失效
		err := contextError(stream.Context())
		if err != nil {
			return err
		}

		//从流中获取请求
		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}

		//从请求中获取电脑ID和评级。
		laptopID := req.GetLaptopId()
		score := req.GetScore()

		log.Printf("received a rate-laptop request: id = %s, score = %.2f", laptopID, score)

		//检查收到的电脑ID是否存在
		found, err := server.laptopStore.Find(laptopID)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot find laptop: %v", err))
		}
		if found == nil {
			return logError(status.Errorf(codes.NotFound, "laptopID %s is not found", laptopID))
		}

		//将新的电脑评级添加到存储，取回更新后的评级对象。
		rating, err := server.ratingStore.Add(laptopID, score)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot add rating to the store: %v", err))
		}

		//创建响应对象。
		res := &pb.RateLaptopResponse{
			LaptopId:     laptopID,
			RatedCount:   rating.Count,
			AverageScore: rating.Sum / float64(rating.Count),
		}

		//将响应发回客户端。
		err = stream.Send(res)
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
	}
	return nil
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline id exceeded"))
	default:
		return nil
	}
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
