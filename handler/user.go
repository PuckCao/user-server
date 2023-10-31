package handler

import (
	"context"
	"fmt"
	pb "github.com/PuckCao/proto/user"
	"time"
	"user-server/dao/db"
)

type UserServer struct{}

func (u *UserServer) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	fmt.Println(req)
	user := db.User{
		Phone:      req.Phone,
		Password:   req.Password,
		Sex:        req.Sex,
		Nickname:   req.Nickname,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	db.RegisterUser(&user)
	return &pb.RegisterResp{
		ErrMsg: &pb.ErrMsg{
			Code: 0,
			Msg:  "success",
		},
	}, nil

}
