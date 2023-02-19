package main

import (
	"douyin-easy/cmd/api/router"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/obs-opentelemetry/tracing"
)

/*
 @Author: 71made
 @Date: 2023/02/19 13:35
 @ProductName: main.go
 @Description: 服务启动函数
*/

func perInit(h *server.Hertz) {
	router.Init(h)

	// hlog init
	// 配置 hertz 日志格式, 默认格式
	hlog.SetLogger(hlog.DefaultLogger())
	hlog.SetLevel(hlog.LevelInfo)
}

func main() {
	tracer, cfg := tracing.NewServerTracer()

	h := server.New(
		server.WithHostPorts(":8080"),
		server.WithMaxRequestBodySize(50*1024*1024),
		//server.WithTransport(standard.NewTransporter),
		tracer,
	)

	perInit(h)

	// use otel mw
	h.Use(tracing.ServerMiddleware(cfg))

	h.Spin()

}
