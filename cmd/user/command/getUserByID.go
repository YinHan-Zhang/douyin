package command

import (
	"context"
	"errors"

	db "douyin-easy/cmd/user/dal"
	pack "douyin-easy/cmd/user/pack"
	usrv "douyin-easy/grpc_gen/user"

	"gorm.io/gorm"
)

type MGetUserService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

// MGetUser get user info by userID
func (s *MGetUserService) MGetUser(req *usrv.QueryUserRequest) (*usrv.User, error) {
	modelUser, err := db.GetUserByID(s.ctx, req.UserId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	user, err := pack.UserMessage(s.ctx, modelUser)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *MGetUserService) MGetUsers(userIDs []int64) ([]*usrv.User, error) {
	modelUsers, err := db.MGetUsers(s.ctx, userIDs)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	users, err := pack.UserMessages(s.ctx, modelUsers)
	if err != nil {
		return nil, err
	}
	return users, nil
}
