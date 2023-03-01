package client

import (
	"bufio"
	"context"
	"fmt"
	"grpctest/pb"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LaptopClient is a client to call laptop service RPCs
type LaptopClient struct {
	service pb.LaptopServiceClient
}

// NewLaptopClient returns a new laptop client
func NewLaptopClient(cc *grpc.ClientConn) *LaptopClient {
	service := pb.NewLaptopServiceClient(cc) //创建一个新的笔记本电脑服务客户端
	return &LaptopClient{service}
}

//创建一个随机的电脑
func (laptopClient *LaptopClient) CreateLaptop(laptop *pb.Laptop) {
	//生成一个新的laptop请求对象
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	//为请求设置超时
	//在goalng中,我们将使用context来做到这一点。
	//该函数返回一个上下文和一个取消对象。
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//使用请求和背景上下文调用notebookClient.Createlaptop()
	res, err := laptopClient.service.CreateLaptop(ctx, req)
	if err != nil {
		//如果发生错误，我们将其转换为状态对象,这样我们就可以检查返回的状码。
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			//如果已经存在,写个简单的日志记录一下就可以了。
			log.Print("laptop already exists")
		} else {
			//否则记录这个严重的错误。
			log.Fatal("can not create laptop:%v", err)
		}
	}

	//全部顺利执行后，我们只需要写一个日志，说明笔记本电脑是用这个ID创建的
	log.Printf("created laptop with id:%s", res.Id)

}

func (laptopClient *LaptopClient) SearchLaptop(filter *pb.Filter) {
	log.Print("search filter:", filter) //写一个日志显示过滤器的值

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //创建一个超时为5秒的上下文
	defer cancel()

	req := &pb.SearchLaptopRequest{Filter: filter}             //使用过滤器创建一个SearchLaptopRequest对象
	stream, err := laptopClient.service.SearchLaptop(ctx, req) //调用此函数获取流文件
	if err != nil {
		log.Fatal("can not search laptop:", err)
	}

	for { //使用for循环从流中获取多个响应
		res, err := stream.Recv()
		if err == io.EOF { //读到流末尾
			return
		}
		if err != nil {
			log.Fatal("can not receive response:", err)
		}

		laptop := res.GetLaptop() //一切顺利的话从流中获取电脑，然后显示信息
		log.Print(" - found:", laptop.GetId())
		log.Print(" + brand:", laptop.GetBrand())
		log.Print(" + name:", laptop.GetName())
		log.Print(" + cpu corse", laptop.GetCpu().GetNumberCores())
		log.Print(" + cpu min ghz", laptop.GetCpu().GetMinGhz())
		log.Print(" + ram:", laptop.GetRam().GetValue(), laptop.GetRam().GetValue())
		log.Print(" + price:", laptop.GetPriceUsd(), "usd")
	}
}
func (laptopClient *LaptopClient) UploadImage(laptopID string, imagePath string) {
	//打开图像文件
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal("cannot open image file:", err)
	}
	defer file.Close()

	//创建超时为5秒的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//使用该上下文调用UploadImage,返回一个流对象
	stream, err := laptopClient.service.UploadImage(ctx)
	if err != nil {
		log.Fatal("cannot upload image file:", err)
	}

	//创建要向服务器发送的请求
	req := &pb.UploadImageRequest{
		Data: &pb.UploadImageRequest_Info{
			Info: &pb.ImageInfo{
				LaptopId:  laptopID,
				ImageType: ".jpg",
			},
		},
	}

	//将请求发送到服务器
	err = stream.Send(req)
	if err != nil {
		log.Fatal("can not sent image info:", err, stream.RecvMsg(nil))
	}

	//创建一个缓冲区读取器，分块读取图像文件的内容
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024) //假设每块是1KB或1024字节

	//在for循环中读取图像块文件
	for {
		n, err := reader.Read(buffer) //将数据读取到缓冲区,返回读取的字节数和错误
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("can not read chunk to buffer:", err)
		}

		//使用数据块创建一个新请求
		req := &pb.UploadImageRequest{
			Data: &pb.UploadImageRequest_ChunkData{
				ChunkData: buffer[:n], //确保块只包含缓冲区的前n个字节
			},
		}

		//发送到服务器
		err = stream.Send(req)
		if err != nil {
			err2 := stream.RecvMsg(nil) //更清楚显示错误
			log.Fatal("can not send chunk to server:", err, err2)
		}

		//接收服务器的响应
		res, err := stream.CloseAndRecv()
		if err != nil {
			log.Fatal("can not receive response:", err)
		}
		log.Print("image upload with id :%s,size: %d", res.GetId(), res.GetSize())
	}
}

func (laptopClient *LaptopClient) RateLaptop(laptopIDs []string, scores []float64) error {
	//定义五秒后超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//得到流
	stream, err := laptopClient.service.RateLaptop(ctx)
	if err != nil {
		return fmt.Errorf("can not rate laptop: %v", err)
	}

	//创建一个通道来接收来自服务器的响应。
	waitResponse := make(chan error)
	//创建goroutine来接收响应
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Print("no more response")
				waitResponse <- nil
				return
			}
			if err != nil {
				waitResponse <- fmt.Errorf("can not receive stream response:%v", err)
				return
			}

			log.Print("received response:", res)
		}
	}()

	//发送请求
	for i, laptopID := range laptopIDs { //遍历电脑列表。
		req := &pb.RateLaptopRequest{ //为每台电脑创建请求
			LaptopId: laptopID,
			Score:    scores[i],
		}

		err := stream.Send(req) //发送到服务端
		if err != nil {
			return fmt.Errorf("can not send stream request:%v-%v", err, stream.RecvMsg(nil))
		}

		log.Print("sent request: %v", req)
	}

	//发送结束后告诉服务器我们不再发送任何数据。
	err = stream.CloseSend()
	if err != nil {
		return fmt.Errorf("can not close send: %v", err)
	}

	err = <-waitResponse
	return err

}
