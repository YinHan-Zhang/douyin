package dal

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

/*
 @Author: 71made
 @Date: 2023/01/24 23:05
 @ProductName: user.go
 @Description: 用户表数据模型
*/

type User struct {
	gorm.Model
	Username       string
	Password       string
	Avatar         string
	FollowCount    int64
	FollowerCount  int64
	TotalFavorited int64
}

func (u *User) TableName() string {
	return "user"
}

func QueryUsers(ctx context.Context, username string) ([]User, error) {
	res := make([]User, 0)
	if err := GetInstance().WithContext(ctx).
		Model(&User{}).Where("username = ?", username).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// QueryUsersByIds
// 通过 ids 查询 model.User, 并按传入 ids 序列进行排序
func QueryUsersByIds(ctx context.Context, userIds []int64) ([]User, error) {
	res := make([]User, 0)
	if len(userIds) == 0 {
		return res, nil
	}
	// 构造排序条件
	str := strings.ReplaceAll(fmt.Sprintf("%v", userIds), " ", ",")
	// 截取中间 id 序列
	str = str[1 : len(str)-1]
	if err := GetInstance().
		Model(&User{}).WithContext(ctx).Where("id in ?", userIds).
		Order("Field(id," + str + ")").
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryUser(ctx context.Context, username string) (*User, error) {
	res, err := QueryUsers(ctx, username)
	if err != nil || len(res) == 0 {
		return nil, err
	}

	return &res[0], nil
}

func QueryUserById(ctx context.Context, userId int64) (*User, error) {
	res := make([]User, 0)
	if err := GetInstance().WithContext(ctx).Where("id = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}
	return &res[0], nil
}

func IsExistUser(ctx context.Context, username string) (bool, error) {
	if len(username) == 0 {
		return false, errors.New("username can't be empty")
	}
	ids := make([]int64, 0)
	if err := GetInstance().WithContext(ctx).
		Model(&User{}).
		Select("id").
		Where("username = ?", username).
		Find(&ids).Error; err != nil {
		return false, err
	}
	return len(ids) != 0, nil
}

func CreateUser(ctx context.Context, user *User) error {
	return GetInstance().WithContext(ctx).Create(user).Error
}
