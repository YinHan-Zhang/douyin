package core

import (
	"context"
	"douyin-easy/cmd/api/biz"
	service "douyin-easy/cmd/api/biz/service/core"
	"douyin-easy/pkg/configs"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"net/http"
	"strconv"
)

/*
 @Author: 71made
 @Date: 2023/02/15 11:41
 @ProductName: user.go
 @Description:
*/

var userService = service.UserServiceImpl()

// UserInfo
// @router /douyin/user/ [GET]
func UserInfo(ctx context.Context, c *app.RequestContext) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		resp := biz.NewFailureResponse("请求参数异常")
		hlog.Errorf(fmt.Sprintf("msg : %s\n error: %v", "user_id 类型转换错误", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	var thisUserId int64
	// 获取 JWT 回设的 userId
	v, _ := c.Get(configs.IdentityKey)
	if v != nil {
		thisUserId = v.(*biz.User).Id
	} else {
		thisUserId = biz.NotLoginUserId
	}

	resp := userService.UserInfo(ctx, userId, thisUserId)
	c.JSON(http.StatusOK, resp)
}

// UserLogin
// @router /douyin/user/login/ [POST]
// 注意: 使用 JWT 中间件后, 路由不再直接绑定此方法, 而是在 JWT 中配置的 Authenticator(认证器) 方法中调用,
// 而路由则直接绑定的 JWT 提供的 LoginHandler 方法. 对于 UserLogin 返回的 resp 和 status 调整为
// 通过 app.RequestContext 传递, 最终再通过 JWT 配置的 Unauthorized(未授权的)、LoginResponse 等回调方法写入 JSON 返回给客户端
func UserLogin(ctx context.Context, c *app.RequestContext) {
	req := &service.UserLoginRequest{}

	err := c.BindAndValidate(req)
	if err != nil {
		hlog.Error(err)
		c.JSON(http.StatusBadRequest, biz.NewErrorResponse(fmt.Errorf("参数绑定失败")))
		return
	}

	resp := userService.Login(ctx, req)
	c.Set("resp", resp)
}

// UserRegister
// @router /douyin/user/register/ [POST]
// 注意: 使用 JWT 中间件后, 不直接返回响应结果, 而是通过后续的 app.HandlerFunc 中
// 去调用 JWT 提供的函数生成 token 一起响应返回
func UserRegister(ctx context.Context, c *app.RequestContext) {
	req := &service.UserRegisterRequest{}

	err := c.BindAndValidate(req)
	if err != nil {
		hlog.Error(err)
		c.Set("status", http.StatusBadRequest)
		c.Set("resp", biz.NewErrorResponse(fmt.Errorf("参数绑定失败")))
		return
	}

	resp := userService.Register(ctx, req)

	c.Set("status", http.StatusOK)
	c.Set("resp", resp)
}
