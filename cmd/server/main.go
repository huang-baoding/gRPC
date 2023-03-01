package main

import (
	"flag"
	"fmt"
	"grpctest/pb"
	"grpctest/service"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	secretKey     = "secret"
	tokenDuration = 15 * time.Minute
)

//为了测试新的登录API，我们必须添加一些种子用户
func seedUsers(userStore service.UserStore) error {
	err := createUser(userStore, "admin1", "secret", "admin") //创建一个管理员
	if err != nil {
		return err
	}
	return createUser(userStore, "user1", "secret", "user") //创建一个用户
}

func createUser(userStore service.UserStore, username, password, role string) error {
	user, err := service.NewUser(username, password, role)
	if err != nil {
		return err
	}
	return userStore.Save(user)
}

func accessibleRoles() map[string][]string {
	const laptopServicePath = "/pb.LaptopService"

	return map[string][]string{
		laptopServicePath + "CreateLaptop": {"admin"},
		laptopServicePath + "UploadImage":  {"admin"},
		laptopServicePath + "RateLaptop":   {"admin", "user"},
	}
}

func main() {
	//使用flag.Int从命令行参数获取端口
	port := flag.Int("port", 0, "the server port")
	//解析标志
	flag.Parse()
	//打印一个简单的日志
	log.Printf("start server on port %d", *port)

	//将身份验证添加到gRPC服务
	userStore := service.NewInMemoryUserStore()
	err := seedUsers(userStore)
	if err != nil {
		log.Fatal("cannot seed users: ", err)
	}

	jwtManager := service.NewJWTManager(secretKey, tokenDuration) //使用密钥和令牌持续时间创建一个新的JWT管理器
	//创建一个新的身份验证服务器
	authServer := service.NewAuthServer(userStore, jwtManager)

	laptopStore := service.NewInMemoryLaptopStore()
	imageStore := service.NewDiskImageStore("./img/") //在img文件夹中保存上传的图像
	//使用内存存储创建一个新的laptop服务器对象
	ratingStore := service.NewInMemoryRatingStore()
	LaptopServer := service.NewLaptopServer(laptopStore, imageStore, ratingStore)

	interceptor := service.NewAuthInterceptor(jwtManager, accessibleRoles())
	//创建一个新的gRPC服务器
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()), //他需要一个一元服务器拦截器函数作为输入
		grpc.StreamInterceptor(interceptor.Stream()),
	)

	pb.RegisterAuthServiceServer(grpcServer, authServer)
	//向gRPC服务器上注册laptop服务器
	pb.RegisterLaptopServiceServer(grpcServer, LaptopServer)
	reflection.Register(grpcServer) //调用反射注册

	// 用之前得到的端口创建一个地址字符串
	address := fmt.Sprintf("0.0.0.0:%d", *port)
	//监听此tcp上的连接
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("can not start server:%v", err)
	}

	//调用grpcServer.Server()来启动服务
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatal("can not start server:%v", err)
	}
}
