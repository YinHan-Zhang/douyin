package rpc

import (
	"douyin-easy/cmd/api/biz"
	rpc "douyin-easy/grpc_gen"
)

/*
 @Author: 71made
 @Date: 2023/02/15 19:34
 @ProductName: resp.go
 @Description:
*/

func NewBizResponse(resp *rpc.BaseResponse) *biz.Response {
	if resp == nil {
		return biz.NewFailureResponse("无效响应")
	}
	return &biz.Response{
		StatusCode: int32(resp.StatusCode),
		StatusMsg:  resp.StatusMsg,
	}
}
