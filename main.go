package main

import (
	"context"
	"fmt"
	pb "github.com/PuckCao/proto/user"
	"google.golang.org/grpc"
	"net"
)

type userServer struct{}

func (u *userServer) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	fmt.Println(req)
	if req.Password == "123456" {
		return &pb.RegisterResp{
			ErrMsg: &pb.ErrMsg{
				Code: 10001,
				Msg:  "password error",
			},
		}, nil
	}
	return &pb.RegisterResp{
		ErrMsg: &pb.ErrMsg{
			Code: 0,
			Msg:  "success",
		},
	}, nil

}

func main() {
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
	pb.RegisterUserServer(s, &userServer{})
	// 4.启动服务端
	s.Serve(listener)
}
