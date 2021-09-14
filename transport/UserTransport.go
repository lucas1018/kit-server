package transport

import (
	"fund/endpoint"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func DecodeUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	// 限定参数来源,我们直接从url的params中获取用户id
    // 请求格式类似于：http://127.0.0.1?uid=101
	if r.URL.Query().Get("uid") != "" {
		uid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
		return endpoint.UserRequest{Uid: int32(uid)}, nil
	}
	return nil, errors.New("参数错误")
}

func EncodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
    // 将返回体body置为json格式
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

