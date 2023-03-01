package main

import (
	"flag"
	"fmt"
	"grpctest/client"
	"grpctest/pb"
	"grpctest/sample"
	"log"
	"strings"
	"time"

	"google.golang.org/grpc"
)

//测试创建电脑rpc的功能。
func testCreateLaptop(laptopClient *client.LaptopClient) {
	laptopClient.CreateLaptop(sample.NewLaptop())
}

//测试搜索电脑rpc
func testSearchLaptop(laptopClient *client.LaptopClient) {
	for i := 0; i < 10; i++ { //创建十个随机电脑
		laptopClient.CreateLaptop(sample.NewLaptop())
	}

	filter := &pb.Filter{ //创建一个新的搜索过滤器
		MaxPriceUsd: 3000, //查找电脑的最大价格
		MinCpuCores: 4,
		MinCpuGhz:   2.5,
		MinRam:      &pb.Memory{Value: 8, Uint: pb.Memory_GIGABYTE}, //最小有8G的内存
	}

	laptopClient.SearchLaptop(filter)
}

//测试上传图片的rpc
func testUploadImage(laptopClient *client.LaptopClient) {
	//首先是生成一个随机的电脑。
	laptop := sample.NewLaptop()
	//在服务器上创建它
	laptopClient.CreateLaptop(laptop)
	//将此图像上传到服务器。
	laptopClient.UploadImage(laptop.GetId(), "tmp/laptop.jpg")
}

func testRateLaptop(laptopClient *client.LaptopClient) {
	n := 3 //假设我们要对三个电脑进行评级。
	laptopIDs := make([]string, n)

	for i := 0; i < n; i++ { //随机生成三个电脑。
		laptop := sample.NewLaptop()
		laptopIDs[i] = laptop.GetId()
		laptopClient.CreateLaptop(laptop)
	}

	//对电脑进行评级。
	scores := make([]float64, n)
	for {
		fmt.Print("rate laptop (y/n)? ")
		var answer string
		fmt.Scan(&answer)

		if strings.ToLower(answer) != "y" {
			break
		}

		for i := 0; i < n; i++ {
			scores[i] = sample.RandomLaptopScore()
		}

		err := laptopClient.RateLaptop(laptopIDs, scores)
		if err != nil {
			log.Fatal(err)
		}
	}

}

const (
	username        = "admin1"
	password        = "secret"
	refreshDuration = 30 * time.Second
)

func authMethods() map[string]bool {
	const laptopServicePath = "/pb.LaptopService"

	return map[string]bool{
		laptopServicePath + "CreateLaptop": true,
		laptopServicePath + "UploadImage":  true,
		laptopServicePath + "RateLaptop":   true,
	}
}

func main() {
	//首先是从命令行参数中获取服务器地址。
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	//写一个简单的日志，说我们正在拨打这个服务器
	log.Printf("dial server %s", *serverAddress)

	//使用输入地址调用grpc.Dial()函数
	cc1, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	authClient := client.NewAuthClient(cc1, username, password)
	interceptor, err := client.NewAuthInterceptor(authClient, authMethods(), refreshDuration)
	if err != nil {
		log.Fatal("cannot create auth interceptor: ", err)
	}

	cc2, err := grpc.Dial(
		*serverAddress,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptor.Unary()),
		grpc.WithStreamInterceptor(interceptor.Stream()),
	)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	//使用连接创建一个新的laptop客户端对象
	laptopClient := client.NewLaptopClient(cc2)

	testRateLaptop(laptopClient)

}
