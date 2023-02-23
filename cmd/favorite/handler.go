package main

import (
	"context"
	"douyin-easy/cmd/favorite/dal"
	"douyin-easy/cmd/favorite/pack"
	fsvr "douyin-easy/grpc_gen/favorite"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

/*
@Author: 71made
@Date: 2023/02/19 13:30
@ProductName: handler.go
@Description: 实现 gRPC 生成代码中提供的服务接口，并提供构建 server 的函数供 main 方法调用
*/

type FavoriteManagementServer struct {
	fsvr.UnimplementedFavoriteManagementServer
}

func newServer(opts ...grpc.ServerOption) *grpc.Server {
	svr := grpc.NewServer(opts...)
	fsvr.RegisterFavoriteManagementServer(svr, &FavoriteManagementServer{})
	return svr
}
func (fms FavoriteManagementServer) Action(ctx context.Context, request *fsvr.FavoriteActionRequest) (*fsvr.FavoriteActionResponse, error) {
	//TODO implement me
	fmt.Println("申请Action服务")
	resp := &fsvr.FavoriteActionResponse{}
	f := &dal.Favorite{
		UserId:  uint(request.UserId),
		VideoId: uint(request.VideoId),
	}
	switch request.Type {
	case fsvr.Action_Commit:
		{
			f.FavoriteType = dal.Favorable
		}
	case fsvr.Action_Cancel:
		{
			f.FavoriteType = dal.Unfavorable
		}
	case fsvr.Action_Unknown:
		fallthrough
	default:
		resp.BaseResponse = pack.NewFailureResponse("参数异常错误, 操作失败")
		return resp, nil
	}
	// 查找点赞记录
	found, err := dal.QueryFavorite(ctx, request.UserId, request.VideoId)
	if err != nil {
		grpclog.Error(err)
		resp.BaseResponse = pack.NewErrorResponse(err)
		return resp, nil
	}

	// 没有记录则创建, 有则更新
	if found == nil {
		err = dal.CreateFavorite(ctx, f)
	} else if found.FavoriteType != f.GetFavoriteType() {
		// 并且只对于不同的 type, 才触发更新
		err = dal.UpdateFavorite(ctx, f)
	}

	if err != nil {
		grpclog.Error(err)
		resp.BaseResponse = pack.NewErrorResponse(err)
		return resp, nil
	}
	resp.BaseResponse = pack.NewSuccessResponse("操作成功")
	return resp, nil
}

func (fms FavoriteManagementServer) QueryFavorites(ctx context.Context, request *fsvr.QueryFavoritesRequest) (*fsvr.QueryFavoritesResponse, error) {
	//TODO implement me
	// 内部抽取处理函数
	fmt.Println("申请 QueryFavorites服务")
	packQueryRes := func(favorites []*dal.Favorite, err error) *fsvr.QueryFavoritesResponse {
		resp := &fsvr.QueryFavoritesResponse{}
		if err != nil {
			grpclog.Error(err)
			resp.BaseResponse = pack.NewErrorResponse(err)
			resp.FavoriteList = make([]*fsvr.Favorite, 0)
			return resp
		}

		resp.FavoriteList = pack.Favorites(favorites)
		resp.BaseResponse = pack.NewSuccessResponse("获取成功")
		return resp
	}

	if request.VideoIds == nil || len(request.VideoIds) == 0 {
		favorites, err := dal.QueryUserFavorites(ctx, request.UserId)
		return packQueryRes(favorites, err), nil
	} else {
		favorites, err := dal.QueryFavorites(ctx, request.UserId, request.VideoIds)
		return packQueryRes(favorites, err), nil
	}

}
func (fms FavoriteManagementServer) mustEmbedUnimplementedFavoriteManagementServer() {
}
