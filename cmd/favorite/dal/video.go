package dal

import (
	"context"
	"douyin-easy/pkg/configs"
	"gorm.io/gorm"
)

/*
 @Author: 71made
 @Date: 2023/01/28 22:17
 @ProductName: video.go
 @Description: 用户视频表数据模型
*/

type Video struct {
	gorm.Model
	AuthorId      uint
	PlayUri       string
	CoverUri      string
	FavoriteCount int64
	CommentCount  int64
	Title         string
}

func (v *Video) TableName() string {
	return configs.VideoTable
}

func QueryUserIdByVideoId(ctx context.Context, videoId uint) (int64, error) {
	video := Video{}
	err := GetInstance().WithContext(ctx).
		Where("id=?", videoId).
		Take(&video).Error
	return int64(video.AuthorId), err
}
