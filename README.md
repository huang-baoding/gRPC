项目介绍
在proto文件中定义需要传递的消息内容（当前项目的proto Message定义为带有各种配置信息的laptop）
客户端生成各种配置的laptop信息以双向流的方式发送到服务端，然后服务端给客户端发回响应


部署运行
clone到本地后执行go mod tidy初始化项目依赖
执行 go run cmd/server/main.go -port 9090 运行服务端
执行 go run cmd/client/main.go -address 0.0.0.0:9090 运行客户端


目录结构
cmd文件夹存放客户端和服务端进行连接和操作的代码
client文件夹存放客户端实现具体功能的函数代码
service文件夹存放服务端实现具体功能的函数代码
proto文件夹存放的是.proto文件
pb文件夹存放的是所有由.proto文件生成的.go文件
serializer文件夹实现的是将laptop对象序列化为文件
tmp文件夹是存放测试serializer文件夹功能后生成的二进制文件和json文件
img文件夹存放的是用于测试客户端流生成的图片文件




具体实现：

(1)双向流：客户端和服务端并行发送多个请求和多个响应。为客户端编写一个API以对分数从1到10的笔记本电脑流进行评分。服务器将以每台每台笔记本电脑的平均分数流进行响应。

(2)用Evans客户端来玩转gRPC的反射(Evans还没安装成功)
gRPC反射是服务器的可选扩展，用于帮助客户端构建请求，而无需事先生成存根，这对于客户端在实际实施之前探索gRPC API非常有用
在golang服务器中使用：
    先导入反射包："google.golang/grpc/reflection"
    然后调用反射注册:reflection.Register(grpcServer)

(3)gRPC拦截器（用来对用户进行身份验证和授权，类似于一个中间件功能，可以在客户端和服务器端添加）
服务端拦截器是gRPC服务器在到达实际RPC方法之前调用的函数，可以用于多种用途，例如：日志记录、跟踪、速率限制、身份验证、授权
客户端拦截器是gRPC客户端在调用实际RPC之前调用的函数
（1）服务端拦截器
    使用JSON Web令牌（JWT）授权访问我们的gRPC API
    只有具有某些特定角色的用户才能调用我们服务器上的特定API
（2）客户端拦截器
    登录用户并将JWT附加到请求中，然后再调用gRPC API

对于服务端有两种类型的拦截器：一种用于一元RPC，另一种用于流RPC.

当前项目只实现到gRPC拦截器

