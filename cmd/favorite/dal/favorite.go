package dal

import (
	"context"
	"douyin-easy/pkg/configs"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
	"time"
)

/*
@Author: 71made
@Date: 2023/02/19 13:29
@ProductName: favorite.go
@Description: 定义 favorite 数据库实体 struct，封装实现 favorite 相关 curd 操作
*/
const (
	Favorable   = 1
	Unfavorable = 2
)

type Favorite struct {
	ID           uint `gorm:"primarykey"`
	UserId       uint
	VideoId      uint
	FavoriteType uint `gorm:"column:is_favorite"` // 1-点赞 2-取消点赞
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (f *Favorite) TableName() string {
	return "favorite"
}

func (f *Favorite) IsFavorite() bool {
	return f.FavoriteType == Favorable
}

func (f *Favorite) GetFavoriteType() uint {
	// 过滤其他参数
	if f.FavoriteType != Favorable {
		f.FavoriteType = Unfavorable
	}
	return f.FavoriteType
}

// BeforeCreate
// 通过 GORM 提供的 Hook 实现关联更新 video 记录的 favorite_count
func (f *Favorite) BeforeCreate(tx *gorm.DB) (err error) {
	err1 := f.syncUpdateFavoriteCount(tx, gorm.Expr("favorite_count + 1"))
	if err1 != nil {
		return err1
	}
	err2 := f.syncUpdateTotalFavoriteCount(tx, gorm.Expr("total_favorited + 1"))
	if err2 != nil {
		return err2
	}
	return nil
}

// BeforeUpdate
// 同理, 通过 Hook 实现关联更新 video 记录的 favorite_count
func (f *Favorite) BeforeUpdate(tx *gorm.DB) (err error) {
	var expr_v clause.Expr
	var expr_u clause.Expr
	if f.IsFavorite() {
		expr_v = gorm.Expr("favorite_count + 1")
		expr_u = gorm.Expr("total_favorited + 1")
	} else {
		expr_v = gorm.Expr("favorite_count - 1")
		expr_u = gorm.Expr("total_favorited - 1")
	}
	err1 := f.syncUpdateFavoriteCount(tx, expr_v)
	if err1 != nil {
		return err1
	}
	err2 := f.syncUpdateTotalFavoriteCount(tx, expr_u)
	if err2 != nil {
		return err2
	}
	return nil
}

func (f *Favorite) syncUpdateFavoriteCount(tx *gorm.DB, expr clause.Expr) (err error) {

	updateRes := tx.Model(&Video{}).Where("id = ?", f.VideoId).
		Update("favorite_count", expr)
	if err = updateRes.Error; err != nil {
		return err
	}
	if updateRes.RowsAffected <= 0 {
		return errors.New("update user_video record fail")
	}
	if updateRes.RowsAffected > 1 {
		// 对于影响数超过 1 的更新, 逻辑上是不合理的, 可能是 video 产生脏数据
		// 实际上, 在主键约束下不可能出现此情况, 仅做兜底处理
		return errors.New("user_video table records is dirty")
	}
	return nil
}

func (f *Favorite) syncUpdateTotalFavoriteCount(tx *gorm.DB, expr clause.Expr) (err error) {
	ctx := tx.Statement.Context
	userId, err := QueryUserIdByVideoId(ctx, f.VideoId)
	updateRes := tx.Table(configs.UserTable).Where("id = ?", userId).
		Update("total_favorited", expr)
	if err = updateRes.Error; err != nil {
		return err
	}
	if updateRes.RowsAffected <= 0 {
		return errors.New("update user record fail")
	}
	if updateRes.RowsAffected > 1 {
		// 对于影响数超过 1 的更新, 逻辑上是不合理的, 可能是 user 产生脏数据
		// 实际上, 在主键约束下不可能出现此情况, 仅做兜底处理
		return errors.New("user table records is dirty")
	}
	return nil
}

func CreateFavorite(ctx context.Context, f *Favorite) error {

	return GetInstance().WithContext(ctx).Create(f).Error
}

func UpdateFavorite(ctx context.Context, f *Favorite) error {
	err := GetInstance().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		updateRes := tx.Model(f).
			Where("user_id", f.UserId).Where("video_id", f.VideoId).
			Update("is_favorite", f.GetFavoriteType())

		if updateRes.Error != nil {
			return updateRes.Error
		}
		if updateRes.RowsAffected <= 0 {
			return errors.New("update favorite record fail")
		}
		if updateRes.RowsAffected > 1 {
			// 做兜底处理
			return errors.New("favorite table records is dirty")
		}
		return nil
	})

	return err
}

func QueryFavorite(ctx context.Context, userId, videoId int64) (*Favorite, error) {
	res := make([]Favorite, 0)
	if err := GetInstance().WithContext(ctx).
		Model(&Favorite{}).
		Where("user_id = ?", userId).
		Where("video_id = ?", videoId).
		Find(&res).Error; err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}
	return &res[0], nil
}
func QueryFavorites(ctx context.Context, userId int64, videoIds []int64) ([]*Favorite, error) {
	res := make([]*Favorite, 0)

	if len(videoIds) == 0 {
		return res, nil
	}

	if err := GetInstance().WithContext(ctx).
		Model(&Favorite{}).
		Where("user_id = ?", userId).
		Where("video_id in ?", videoIds).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
func QueryFavoriteVideos(ctx context.Context, userId int64) ([]Video, error) {
	videoIds := make([]int64, 0)
	if err := GetInstance().WithContext(ctx).
		Select("video_id").Model(&Favorite{}).
		Where("user_id = ?", userId).
		Where("is_favorite = ?", Favorable).
		Order("updated_at DESC").
		Find(&videoIds).Error; err != nil {
		return nil, err
	}

	if len(videoIds) == 0 {
		return make([]Video, 0), nil
	}

	res := make([]Video, len(videoIds))
	// 构造排序条件
	str := strings.ReplaceAll(fmt.Sprintf("%v", videoIds), " ", ",")
	// 截取中间 id 序列
	str = str[1 : len(str)-1]
	if err := GetInstance().WithContext(ctx).
		Model(&Video{}).Where("id in ?", videoIds).
		Order("Field(id," + str + ")").
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
func QueryUserFavorites(ctx context.Context, userId int64) ([]*Favorite, error) {
	res := make([]*Favorite, 0)
	if err := GetInstance().WithContext(ctx).
		Model(&Favorite{}).
		Where("user_id = ?", userId).
		Where("is_favorite = ?", Favorable).
		Order("updated_at DESC").
		Find(&res).Error; err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return make([]*Favorite, 0), nil
	}

	return res, nil
}
