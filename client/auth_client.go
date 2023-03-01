package client

import (
	"context"
	"grpctest/pb"
	"time"

	"google.golang.org/grpc"
)

// 为了调用身份验证RPC服务
type AuthClient struct {
	service  pb.AuthServiceClient //服务
	username string               //同于登录验证的账号密码
	password string
}

// NewAuthClient returns a new auth client
func NewAuthClient(cc *grpc.ClientConn, username string, password string) *AuthClient {
	service := pb.NewAuthServiceClient(cc)
	return &AuthClient{service, username, password}
}

//Login函数调用LoginRPC来获取访问令牌
func (client *AuthClient) Login() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.LoginRequest{
		Username: client.username,
		Password: client.password,
	}

	res, err := client.service.Login(ctx, req)
	if err != nil {
		return "", err
	}

	return res.GetAccessToken(), nil
}
