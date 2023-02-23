package pack

import (
	"context"
	db "douyin-easy/cmd/user/dal"
	usrv "douyin-easy/grpc_gen/user"
	"errors"
)

/*
 @Author: 71made
 @Date: 2023/02/19 13:37
 @ProductName: user.go
 @Description: 对 user 数据库实体进行转化，转化成 server 返回类型实体（参考之前的 easy-note 项目）
*/

// UserMessage pack user info
// db.User结构体是用于和数据库交互的
// usrv.User结构体是用于RPC传输信息的，在 user.pb.go文件
func UserMessage(ctx context.Context, u *db.User) (*usrv.User, error) {

	if u == nil {
		return nil, errors.New("db.user is null")
	}

	return &usrv.User{
		Id:              int64(u.ID),
		Name:            u.Username,
		Avatar:          u.Avatar,
		VideoCount:      u.VideoCount,
		FavoriteCount:   u.FavoriteCount,
		FollowCount:     int64(u.FollowerCount),
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
	}, nil
}

// UserMessages pack list of user info
func UserMessages(ctx context.Context, us []*db.User) ([]*usrv.User, error) {
	users := make([]*usrv.User, len(us))
	for _, u := range us {
		user2, err := UserMessage(ctx, u)
		if err != nil {
			return nil, err
		}

		if user2 != nil {
			users = append(users, user2)
		}
	}
	return users, nil
}
