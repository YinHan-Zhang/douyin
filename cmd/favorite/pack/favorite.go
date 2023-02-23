package pack

import (
	"douyin-easy/cmd/favorite/dal"
	fsvr "douyin-easy/grpc_gen/favorite"
)

// import (
//
//	"context"
//	"douyin-easy/cmd/favorite/dal"
//	"douyin-easy/cmd/favorite/message"
//	"github.com/cloudwego/hertz/pkg/common/hlog"
//
// )
//
// /*
// @Author: 71made
// @Date: 2023/02/19 13:37
// @ProductName: favorite.go
// @Description: 对 favorite 数据库实体进行转化，转化成 server 返回类型实体（参考之前的 easy-note 项目）
// */
func Favorite(f *dal.Favorite) *fsvr.Favorite {
	if f == nil {
		return nil
	}

	return &fsvr.Favorite{
		Id:         int64(f.ID),
		UserId:     int64(f.UserId),
		VideoId:    int64(f.VideoId),
		IsFavorite: f.IsFavorite(),
	}
}

func Favorites(fs []*dal.Favorite) []*fsvr.Favorite {
	if fs == nil || len(fs) == 0 {
		return make([]*fsvr.Favorite, 0)
	}

	res := make([]*fsvr.Favorite, len(fs))
	for i, f := range fs {
		res[i] = Favorite(f)
	}

	return res
}

//type FavoriteServiceImpl struct {
//}
//
//func (fs FavoriteServiceImpl) Action(ctx context.Context, request *message.DouyinFavoriteActionRequest) (*message.DouyinFavoriteActionResponse, error) {
//	//TODO implement me
//	// 构建实体
//	f := &dal.Favorite{
//		UserId:       uint(*request.VideoId),
//		VideoId:      uint(*request.VideoId),
//		FavoriteType: uint(*request.ActionType),
//	}
//	// 查找点赞记录
//	found, err := dal.QueryFavorite(ctx, *request.VideoId, *request.VideoId)
//	var resp *message.DouyinFavoriteActionResponse
//	if err != nil {
//		hlog.Error(err)
//		num := int64(1)
//		str := err.Error()
//		*resp = message.DouyinFavoriteActionResponse{StatusCode: &num, StatusMsg: &str}
//		return resp, nil
//	}
//	// 没有记录则创建, 有则更新
//	if found == nil {
//		err = dal.CreateFavorite(ctx, f)
//	} else if found.FavoriteType != f.GetFavoriteType() {
//		// 并且只对于不同的 type, 才触发更新
//		err = dal.UpdateFavorite(ctx, f)
//	}
//	if err != nil {
//		hlog.Error(err)
//		num := int64(1)
//		str := err.Error()
//		*resp = message.DouyinFavoriteActionResponse{StatusCode: &num, StatusMsg: &str}
//		return resp, nil
//	}
//	num := int64(0)
//	str := "操作成功"
//	*resp = message.DouyinFavoriteActionResponse{StatusCode: &num, StatusMsg: &str}
//	return resp, nil
//}
