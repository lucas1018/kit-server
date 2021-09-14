package service

import (
	"context"
	pb "fund/proto"
//	"fmt"
)


type UserService struct {
}

func (s UserService) GetName(ctx context.Context, req *pb.GetNameReq) (resp *pb.GetNameResp, err error) {
	resp = new(pb.GetNameResp)
	if req.UserId == 101 {
		resp.Name = "funder"
		return resp, nil
	}
	resp.Name = "guest"
	return resp, nil
}

