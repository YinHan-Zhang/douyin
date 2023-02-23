package dal

import (
	"context"

	"gorm.io/gorm"
)

/*
 @Author: 71made
 @Date: 2023/02/19 13:29
 @ProductName: user.go
 @Description: 定义 user 数据库实体 struct，封装实现 user 相关 curd 操作
*/

type User struct {
	gorm.Model
	Username        string `gorm:"index:idx_username,unique;type:varchar(40);not null" json:"username"`
	Password        string `gorm:"type:varchar(256);not null" json:"password"`
	VideoCount      int64  `gorm:"default:0" json:"video_count"` //作品数量
	FollowCount     int    `gorm:"default:0" json:"follow_count"`
	FollowerCount   int    `gorm:"default:0" json:"follower_count"`
	FavoriteCount   int64  `gorm:"default:0" json:"favorite_count"` //点赞数量
	Avatar          string `gorm:"default:static/avatar/empty_avatar.jpeg" json:"avatar"`
	BackgroundImage string `gorm:"default:" json:"background_image"` //用户个人页顶部大图
	Signature       string `gorm:"default:帅哥/美女一枚" json:"signature"` //个人简介
}

func (User) TableName() string {
	return "user"
}

// 获取多个用户
func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// 通过用户id获取用户
func GetUserByID(ctx context.Context, userID int64) (*User, error) {
	res := new(User)

	if err := DB.WithContext(ctx).First(&res, userID).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// 注册用户
func CreateUser(ctx context.Context, user *User) (*User, error) {
	result := DB.WithContext(ctx).Create(user)
	return user, result.Error
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
