package dal

import (
	"context"
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
	return "user_video"
}

func QueryVideos(ctx context.Context, options ...PageOption) ([]Video, error) {
	res := make([]Video, 0)

	page := DefaultPage()
	for _, opt := range options {
		opt(page)
	}

	if err := page.Exec(
		GetInstance().WithContext(ctx).
			Order("created_at DESC")).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryVideosByUserId(ctx context.Context, userId int64) ([]Video, error) {
	res := make([]Video, 0)
	if err := GetInstance().WithContext(ctx).
		Where("author_id = ?", userId).Order("created_at DESC").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func IsExistVideo(ctx context.Context, userId int64, title string) (bool, error) {
	ids := make([]int64, 0)
	if err := GetInstance().WithContext(ctx).Model(&Video{}).
		Select("id").
		Where("author_id = ? ", userId).Where("title = ?", title).
		Find(&ids).Error; err != nil {
		return false, err
	}
	return len(ids) != 0, nil
}

func CreateVideo(ctx context.Context, video *Video) error {

	return GetInstance().WithContext(ctx).Create(video).Error
}

func UpdateVideoCoverUri(ctx context.Context, video *Video) error {
	return GetInstance().WithContext(ctx).Model(video).Update("cover_uri", video.CoverUri).Error
}

func QueryUserIdByVideoId(ctx context.Context, videoId uint) (int64, error) {
	video := Video{}
	err := GetInstance().WithContext(ctx).
		Where("id=?", videoId).
		Take(&video).Error
	return int64(video.AuthorId), err
}
