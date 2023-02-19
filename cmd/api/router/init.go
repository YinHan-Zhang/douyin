package router

import (
	"douyin-easy/cmd/api/router/jwt"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 @Author: 71made
 @Date: 2023/02/15 11:43
 @ProductName: init.go
 @Description:
*/

func Init(h *server.Hertz) {
	jwt.Init()
	register(h)
}
