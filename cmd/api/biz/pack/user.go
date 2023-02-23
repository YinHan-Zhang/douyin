package pack

import (
	"douyin-easy/cmd/api/biz"
	usvr "douyin-easy/grpc_gen/user"
	"douyin-easy/pkg/configs"
)

/*
 @Author: 71made
 @Date: 2023/02/21 20:56
 @ProductName: user.go
 @Description:
*/

func BizBaseUser(u *usvr.User, isFollow bool) *biz.BaseUser {
	if u == nil {
		return nil
	}

	return &biz.BaseUser{
		Id:              u.Id,
		Name:            u.Name,
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
		AvatarURL:       configs.ServerAddr + configs.AvatarURIPrefix + u.Avatar,
		IsFollow:        isFollow,
	}
}

func BizUser(u *usvr.User, isFollow bool) *biz.User {
	if u == nil {
		return nil
	}

	return &biz.User{
		BaseUser:           BizBaseUser(u, isFollow),
		WorkCount:          u.VideoCount,
		FavoriteCount:      u.FavoriteCount,
		TotalFavoriteCount: u.TotalFavoriteCount,
		FollowCount:        u.FollowCount,
		FollowerCount:      u.FollowerCount,
	}
}
