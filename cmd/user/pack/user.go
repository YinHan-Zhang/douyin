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

// usrv.User结构体

// type User struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	Id              int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                 // 用户id
// 	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                              // 用户名称
// 	Avatar          string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`                                          // 用户头像
// 	VideoCount      int64  `protobuf:"varint,4,opt,name=video_count,json=videoCount,proto3" json:"video_count,omitempty"`               // 用户视频发布数
// 	FavoriteCount   int64  `protobuf:"varint,5,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty"`      // 用户点赞喜欢视频数
// 	FollowCount     int64  `protobuf:"varint,6,opt,name=follow_count,json=followCount,proto3" json:"follow_count,omitempty"`            // 关注总数
// 	BackgroundImage string `protobuf:"bytes,7,opt,name=background_image,json=backgroundImage,proto3" json:"background_image,omitempty"` //用户个人页顶部大图
// 	Signature       string `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`                                    //个人简介
// 	TotalFavorited  int64  `protobuf:"varint,9,opt,name=total_favorited,json=totalFavorited,proto3" json:"total_favorited,omitempty"`   //获赞数量
// 	IsFollow        bool   `protobuf:"varint,10,opt,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty"`                    // true-已关注，false-未关注
// }

//   }

// UserMessage pack user info
// db.User结构体是用于和数据库交互的
// usrv.User结构体是用于RPC传输信息的，在user.pb.go文件
func UserMessage(ctx context.Context, u *db.User) (*usrv.User, error) {

	if u == nil {
		return nil, errors.New("db.user is null")
	}

	return &usrv.User{
		Id:              int64(u.ID),
		Name:            u.Username,
		Avatar:          u.Avatar,
		VideoCount:      int64(u.VideoCount),
		FavoriteCount:   int64(u.FavoriteCount),
		FollowCount:     int64(u.FollowerCount),
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
	}, nil
}

// UserMessages pack list of user info
func UserMessages(ctx context.Context, us []*db.User) ([]*usrv.User, error) {
	users := make([]*usrv.User, 0)
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
