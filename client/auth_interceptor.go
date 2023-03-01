//构建客户端拦截器,在调用服务器之前，我们将拦截所有的gRPC请求，并为他们附加访问令牌（如果有必要）
package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// AuthInterceptor is a client interceptor for authentication
type AuthInterceptor struct {			//拦截器
	authClient  *AuthClient				//用于登录用户的身份验证客户端对象
	authMethods map[string]bool			//map告诉我们哪个方法需要身份验证的映射
	accessToken string					//最新获取的访问令牌
}

// NewAuthInterceptor returns a new auth interceptor
func NewAuthInterceptor(
	authClient *AuthClient,
	authMethods map[string]bool,
	refreshDuration time.Duration,				//刷新令牌持续时间的参数,告诉我们应该多久调用一次登录API来获取新令牌
) (*AuthInterceptor, error) {
	interceptor := &AuthInterceptor{
		authClient:  authClient,
		authMethods: authMethods,
	}

	err := interceptor.scheduleRefreshToken(refreshDuration)		//刷新令牌
	if err != nil {
		return nil, err
	}

	return interceptor, nil
}

//刷新令牌时间
func (interceptor *AuthInterceptor) scheduleRefreshToken(refreshDuration time.Duration) error {
	err := interceptor.refreshToken()		
	if err != nil {
		return err
	}

	go func() {
		wait := refreshDuration
		for {
			time.Sleep(wait)
			err := interceptor.refreshToken()	
			if err != nil {
				wait = time.Second
			} else {
				wait = refreshDuration
			}
		}
	}()

	return nil
}

//添加拦截器以将令牌附加到请求上下文
// Unary返回一个gRPC一元客户端拦截器
func (interceptor *AuthInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		log.Printf("--> unary interceptor: %s", method)

		if interceptor.authMethods[method] {		//检查此方法是否需要身份验证
			//需要的话，我们必须在调用实际的RPC之前将访问令牌附加到上下文
			return invoker(interceptor.attachToken(ctx), method, req, reply, cc, opts...)
		}

		return invoker(ctx, method, req, reply, cc, opts...) //不需要的话就使用原始上下文调用RPC
	}
}

//Stream和上面的Unary功能差不多
// Stream returns a client interceptor to authenticate stream RPC
func (interceptor *AuthInterceptor) Stream() grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		log.Printf("--> stream interceptor: %s", method)

		if interceptor.authMethods[method] {
			return streamer(interceptor.attachToken(ctx), desc, cc, method, opts...)
		}

		return streamer(ctx, desc, cc, method, opts...)
	}
}

//将令牌附加到输入上下文
func (interceptor *AuthInterceptor) attachToken(ctx context.Context) context.Context {
	//传入输入上下文，以及授权密钥，访问令牌值
	return metadata.AppendToOutgoingContext(ctx, "authorization", interceptor.accessToken)
}


func (interceptor *AuthInterceptor) refreshToken() error {
	accessToken, err := interceptor.authClient.Login()		//刷新令牌而不进行调度
	if err != nil {
		return err
	}

	interceptor.accessToken = accessToken			//返回令牌后将其存储在interceptor.accessToken字段中
	log.Printf("token refreshed: %v", accessToken)

	return nil
}


