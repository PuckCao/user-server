package main

import (
	"fmt"
	pb "github.com/PuckCao/proto/user"
	"google.golang.org/grpc"
	"net"
	"user-server/dao/db"
	"user-server/handler"
)

func main() {
	db.InitMysql()
	// 地址
	addr := "127.0.0.1:11451"
	// 1.监听
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("监听异常:%s\n", err)
	}
	fmt.Printf("监听端口：%s\n", addr)
	// 2.实例化gRPC
	s := grpc.NewServer()
	// 3.在gRPC上注册微服务
	pb.RegisterUserServer(s, &handler.UserServer{})
	// 4.启动服务端
	s.Serve(listener)
}
