package rpc

import (
	"context"
	"douyin-easy/cmd/api/biz"
	rpc "douyin-easy/grpc_gen"
	fsvr "douyin-easy/grpc_gen/favorite"
	"douyin-easy/pkg/configs"
	"douyin-easy/pkg/utils/grpc"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

/*
 @Author: 71made
 @Date: 2023/02/21 14:29
 @ProductName: favorite.go
 @Description:
*/

func favoriteManagementClient() (client fsvr.FavoriteManagementClient, err error) {
	conn, err := grpc.InitClientConn(configs.Etcd, configs.FavoriteServerName)
	if err != nil {
		return nil, err
	}
	return fsvr.NewFavoriteManagementClient(conn), nil
}

func FavoriteAction(ctx context.Context, userId, videoId int64, actionType uint) *biz.Response {
	req := &fsvr.FavoriteActionRequest{
		UserId:  userId,
		VideoId: videoId,
		Type:    fsvr.Action(actionType),
	}

	client, err := favoriteManagementClient()
	if err != nil {
		hlog.Error(err)
		return biz.NewErrorResponse(err)
	}

	resp, err := client.Action(ctx, req)
	if err != nil {
		hlog.Error(err)
		return biz.NewErrorResponse(err)
	}

	return NewBizResponse(resp.BaseResponse)
}

func QueryFavorites(ctx context.Context, userId int64, videoIds []int64) ([]*fsvr.Favorite, *biz.Response) {
	req := &fsvr.QueryFavoritesRequest{
		UserId:   userId,
		VideoIds: videoIds,
	}

	client, err := favoriteManagementClient()
	if err != nil {
		hlog.Error(err)
		return make([]*fsvr.Favorite, 0), biz.NewErrorResponse(err)
	}

	resp, err := client.QueryFavorites(ctx, req)
	if err != nil {
		hlog.Error(err)
		return make([]*fsvr.Favorite, 0), biz.NewErrorResponse(err)
	}

	if resp.BaseResponse != nil && resp.BaseResponse.StatusCode != rpc.Status_OK {
		return make([]*fsvr.Favorite, 0), NewBizResponse(resp.BaseResponse)
	}

	return resp.FavoriteList, NewBizResponse(resp.BaseResponse)

}
