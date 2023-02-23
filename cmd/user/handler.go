package main

import (
	"context"
	"douyin-easy/cmd/user/command"
	"douyin-easy/cmd/user/pack"
	usvr "douyin-easy/grpc_gen/user"
	errself "douyin-easy/pkg/errself"
	"errors"

	"google.golang.org/grpc"
)

/*
 @Author: 71made
 @Date: 2023/02/19 13:30
 @ProductName: handler.go
 @Description: 实现 gRPC 生成代码中提供的服务接口，并提供构建 server 的函数供 main 方法调用
*/

func newServer(opts ...grpc.ServerOption) *grpc.Server {
	svr := grpc.NewServer(opts...)
	usvr.RegisterUserManagementServer(svr, &UserManagementServerImpl{})
	return svr
}

type UserManagementServerImpl struct {
	usvr.UnimplementedUserManagementServer
}

// 验证登录信息是否正确
func (uss UserManagementServerImpl) CheckLoginUser(ctx context.Context, req *usvr.CheckLoginUserRequest) (*usvr.CheckLoginUserResponse, error) {
	resp := &usvr.CheckLoginUserResponse{}

	//判断登录信息是否为空
	if len(req.Password) == 0 || len(req.Username) == 0 {
		err := errors.New("login is null")
		resp.BaseResponse = errself.NewErrorResponse(err)
		return resp, err
	}

	//登录,检查password加密后的值是否与数据库相同
	var checkin = &usvr.CheckLoginUserRequest{
		Username: req.Username,
		Password: req.Password,
	}
	uid, err := command.NewCheckUserService(ctx).CheckUser(checkin)
	if err != nil {
		resp.BaseResponse = errself.NewErrorResponse(err)
		return resp, err
	}
	//返回uid
	resp.UserId = uid
	resp.BaseResponse = errself.NewSuccessResponse("check success")
	return resp, nil
}

func (uss UserManagementServerImpl) CreateUser(ctx context.Context, req *usvr.CreateUserRequest) (*usvr.CreateUserResponse, error) {
	resp := &usvr.CreateUserResponse{}
	user, err := command.NewCreateUserService(ctx).CreateUser(req, Argon2Config)
	if err != nil {
		resp.BaseResponse = errself.NewErrorResponse(err)
		return resp, err
	}
	resp.User, err = pack.UserMessage(ctx, user)
	if err != nil {
		resp.BaseResponse = errself.NewErrorResponse(err)
		return resp, err
	}
	resp.BaseResponse = errself.NewSuccessResponse("create successful!")
	return resp, nil
}

func (uss UserManagementServerImpl) QueryUsers(ctx context.Context, req *usvr.QueryUsersRequest) (*usvr.QueryUsersResponse, error) {
	resp := &usvr.QueryUsersResponse{}
	users, err := command.NewMGetUserService(ctx).MGetUsers(req.UserIds)
	if err != nil {
		resp.BaseResponse = errself.NewErrorResponse(err)
		return resp, err
	}
	resp.UserList = users
	resp.BaseResponse = errself.NewSuccessResponse("find successful!")

	return resp, nil
}
func (uss UserManagementServerImpl) QueryUser(ctx context.Context, req *usvr.QueryUserRequest) (*usvr.QueryUserResponse, error) {
	resp := &usvr.QueryUserResponse{}
	user, err := command.NewMGetUserService(ctx).MGetUser(req)
	if err != nil {
		resp.BaseResponse = errself.NewErrorResponse(err)
		return resp, err
	}
	resp.User = user
	resp.BaseResponse = errself.NewSuccessResponse("find successful!")
	return resp, nil
}
