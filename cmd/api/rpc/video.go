package rpc

import (
	"context"
	"douyin-easy/cmd/api/biz"
	rpc "douyin-easy/grpc_gen"
	vsvr "douyin-easy/grpc_gen/video"
	"douyin-easy/pkg/configs"
	"douyin-easy/pkg/utils/grpc"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

/*
 @Author: 71made
 @Date: 2023/02/17 22:34
 @ProductName: video.go
 @Description:
*/

func videoManagementClient() (client vsvr.VideoManagementClient, err error) {
	conn, err := grpc.InitClientConn(configs.Etcd, configs.VideoServerName)
	if err != nil {
		return nil, err
	}
	return vsvr.NewVideoManagementClient(conn), nil
}

func CreateVideo(ctx context.Context, finalName string, authorId int64, title string) *biz.Response {
	req := &vsvr.CreateVideoRequest{
		VideoFinalName: finalName,
		Title:          title,
		AuthorId:       authorId,
	}

	client, err := videoManagementClient()
	if err != nil {
		hlog.Error(err)
		return biz.NewErrorResponse(err)
	}

	resp, err := client.CreateVideo(ctx, req)
	if err != nil {
		hlog.Error(err)
		return biz.NewErrorResponse(err)
	}

	return NewBizResponse(resp.BaseResponse)
}

func QueryVideosByUserId(ctx context.Context, userId int64) ([]*vsvr.Video, *biz.Response) {

	req := &vsvr.QueryVideosRequest{
		UserId: &userId,
	}

	client, err := videoManagementClient()
	if err != nil {
		hlog.Error(err)
		return make([]*vsvr.Video, 0), biz.NewErrorResponse(err)
	}

	resp, err := client.QueryVideos(ctx, req)
	if err != nil {
		hlog.Error(err)
		return make([]*vsvr.Video, 0), biz.NewErrorResponse(err)
	}

	if resp.BaseResponse != nil && resp.BaseResponse.StatusCode != rpc.Status_OK {
		return make([]*vsvr.Video, 0), NewBizResponse(resp.BaseResponse)
	}

	return resp.VideoList, NewBizResponse(resp.BaseResponse)
}

func QueryFeedVideos(ctx context.Context, limit int, lastTime int64) ([]*vsvr.Video, *biz.Response) {
	req := &vsvr.QueryFeedVideosRequest{
		Limit:    int32(limit),
		LastTime: lastTime,
	}

	client, err := videoManagementClient()
	if err != nil {
		hlog.Error(err)
		return make([]*vsvr.Video, 0), biz.NewErrorResponse(err)
	}

	resp, err := client.QueryFeedVideos(ctx, req)
	if err != nil {
		hlog.Error(err)
		return make([]*vsvr.Video, 0), biz.NewErrorResponse(err)
	}

	if resp.BaseResponse != nil && resp.BaseResponse.StatusCode != rpc.Status_OK {
		return make([]*vsvr.Video, 0), NewBizResponse(resp.BaseResponse)
	}

	return resp.VideoList, NewBizResponse(resp.BaseResponse)
}

func QueryFavoriteVideos(ctx context.Context, userId int64) ([]*vsvr.Video, *biz.Response) {
	req := &vsvr.QueryFavoriteVideosRequest{
		UserId: userId,
	}

	client, err := videoManagementClient()
	if err != nil {
		hlog.Error(err)
		return make([]*vsvr.Video, 0), biz.NewErrorResponse(err)
	}

	resp, err := client.QueryFavoriteVideos(ctx, req)
	if err != nil {
		hlog.Error(err)
		return make([]*vsvr.Video, 0), biz.NewErrorResponse(err)
	}

	if resp.BaseResponse != nil && resp.BaseResponse.StatusCode != rpc.Status_OK {
		return make([]*vsvr.Video, 0), NewBizResponse(resp.BaseResponse)
	}

	return resp.VideoList, NewBizResponse(resp.BaseResponse)
}
