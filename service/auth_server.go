//为客户端提供一个新的服务来登录并获取访问令牌
package service

import (
	"context"
	"grpctest/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthServer is the server for authentication
type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	userStore  UserStore   //用户存储
	jwtManager *JWTManager //JWT管理器
}

// NewAuthServer returns a new auth server
func NewAuthServer(userStore UserStore, jwtManager *JWTManager) pb.AuthServiceServer {
	return &AuthServer{userStore: userStore, jwtManager: jwtManager}
}

// Login is a unary RPC to login user
func (server *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := server.userStore.Find(req.GetUsername()) 		//通过用户名查找用户
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := server.jwtManager.Generate(user)		//找到用户就生成一个新的访问令牌
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	//用生成的访问令牌创建一个新的登录响应对象
	res := &pb.LoginResponse{AccessToken: token}
	return res, nil			//将其返回给客户端
}

