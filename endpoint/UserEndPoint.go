package endpoint

import (
	"fund/service"
	pb "fund/proto"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type UserRequest struct {
	Uid int32 `json:"uid"`
}

type UserResponse struct {
	Result *pb.GetNameResp `json:"result"`
}

func GetUserEndPoint(userService service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserRequest)
		result, nil := userService.GetName(ctx, &pb.GetNameReq{
			UserId: r.Uid,
		})
		return UserResponse{Result: result}, nil
	}
}

