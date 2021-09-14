package main

import (
	"fund/endpoint"
	"fund/service"
	"fund/transport"
	httpTransport "github.com/go-kit/kit/transport/http"
	"net/http"
	"google.golang.org/grpc"
	"flag"
	"fmt"
	"net"
	pb "fund/proto"
)

func main() {
	var (
//		serviceHost = flag.String("service.host", "localhost", "service ip address")
		servicePort = flag.String("service.port", "9001", "service port")
 
		grpcAddr    = flag.String("grpc", ":8001", "gRPC listen address.")
	)
	errChan := make(chan error)

	user := service.UserService{}
	endPoint := endpoint.GetUserEndPoint(user)
	// 构造服务，实现http.handler并包装endpoint层
	serverHandler := httpTransport.NewServer(endPoint, transport.DecodeUserRequest, transport.EncodeUserResponse)
	// http 监听端口，并且使用serverHandler处理随之而来的请求
	go func() {
		fmt.Println("Http Server start at port:" + *servicePort)
		//启动前执行注册
		//registar.Register()
 
		errChan <- http.ListenAndServe(":"+*servicePort, serverHandler)
	}()
	// grpc server
	go func() {
		fmt.Println("grpc Server start at port" + *grpcAddr)
		listener, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			errChan <- err
			return
		}
		baseServer := grpc.NewServer()
		pb.RegisterUserServer(baseServer, user)
		errChan <- baseServer.Serve(listener)
 
	}()

	error := <-errChan
	//服务退出取消注册
	//registar.Deregister()
	fmt.Println(error)
}

