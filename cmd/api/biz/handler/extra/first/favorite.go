package first

import (
	"context"
	"douyin-easy/cmd/api/biz"
	"douyin-easy/cmd/api/biz/service/extra/first"
	"douyin-easy/pkg/configs"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"net/http"
	"strconv"
)

/*
 @Author: 71made
 @Date: 2023/02/21 14:36
 @ProductName: favorite.go
 @Description:
*/

var favoriteServiceImpl = first.FavoriteServiceImpl()

// FavoriteAction
// @router /douyin/favorite/action/ [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	req := &first.FavoriteActionRequest{}

	err := c.BindAndValidate(req)
	if err != nil {
		hlog.Error(err)
		c.JSON(http.StatusBadRequest, biz.NewErrorResponse(fmt.Errorf("参数绑定错误")))
		return
	}

	// 获取 JWT 回设的 userId
	v, _ := c.Get(configs.IdentityKey)
	userId := v.(*biz.User).Id

	req.UserId = userId

	resp := favoriteServiceImpl.Action(ctx, req)
	c.JSON(http.StatusOK, resp)

}

// GetFavoriteList
// @route /douyin/favorite/list/ [GET]
func GetFavoriteList(ctx context.Context, c *app.RequestContext) {

	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		hlog.Error(err)
		c.JSON(http.StatusBadRequest, biz.NewErrorResponse(fmt.Errorf("参数类型转换错误")))
		return
	}

	resp := favoriteServiceImpl.FavoriteList(ctx, userId)
	c.JSON(http.StatusOK, resp)
}
