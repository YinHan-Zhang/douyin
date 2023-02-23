package main

import (
	"douyin-easy/cmd/favorite/dal"
	_ "douyin-easy/cmd/favorite/dal"
	"douyin-easy/pkg/utils"
	"douyin-easy/pkg/utils/etcd"
	"flag"
	"fmt"
	_ "go.etcd.io/etcd/client/v3"
	_ "go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc/grpclog"
	"net"
)

/*
@Author: 71made
@Date: 2023/02/19 13:35
@ProductName: main.go
@Description: 服务启动函数
*/

// Favorite RPC Server 端运行
func main() {

	dal.InitDB()
	//fmt.Println("main started...")
	//
	////注册服务
	//etcdClient, err := etcd.Register("douyin/FavoriteServer", "127.0.0.1:8083")
	//if err != nil {
	//	grpclog.Fatal(err)
	//	fmt.Println(err)
	//}
	//utils.DealSignal(func() {
	//	// 注销注册
	//	_ = etcd.Unregister(etcdClient, "douyin/FavoriteServer", "127.0.0.1:8083")
	//})
	////log.Println("server started...")
	//fmt.Println("server started...")
	////开启监听
	//listen, err := net.Listen("tcp", "127.0.0.1:8083")
	//if err != nil {
	//	panic(err)
	//}
	//svr := newServer()
	////grpclog.Info("Running user grpc server...")
	//fmt.Println("Running user grpc server...")
	//err = svr.Serve(listen)
	//if err != nil {
	//	grpclog.Fatal("User grpc server start failed: ", err)
	//	return
	var port int
	flag.IntVar(&port, "port", 8083, "port")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", "127.0.0.1", port)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	etcdClient, err := etcd.Register("douyin/favorite", addr)
	if err != nil {
		grpclog.Fatal("Favorite grpc server register to ETCD failed: ", err)
	}

	utils.DealSignal(func() {
		// 注销注册
		_ = etcd.Unregister(etcdClient, "douyin/favorite", addr)
	})

	svr := newServer()
	grpclog.Info("Running favorite grpc server...")
	err = svr.Serve(listen)
	if err != nil {
		grpclog.Fatal("Favorite grpc server start failed: ", err)
		return
	}

}

//func main() {
//	dal.InitDB()
//	server := grpc.NewServer()
//
//	// 将FavoriteManagementServer对象注册到server中
//	favorite.RegisterFavoriteManagementServer(server, new(FavoriteManagementServer))
//
//	lis, err := net.Listen("tcp", ":8080")
//	if err != nil {
//		panic(err.Error())
//	}
//	// 提供RPC服务
//	server.Serve(lis)

//tracer, cfg := tracing.NewServerTracer()
//
//h := server.New(
//	server.WithHostPorts(":8080"),
//	server.WithMaxRequestBodySize(50*1024*1024),
//	//server.WithTransport(standard.NewTransporter),
//	tracer,
//)
//
//dal.InitDB()
////minio.Init(ctx)
////jwt.Init()
////router.Init(h)
//root := h.Group("/douyin")
//_favorite := root.Group("/favorite", func(ctx context.Context, c *app.RequestContext) {
//	// 对于 /list/  接口, 在用户未登陆时也可以请求查看其他用户的喜欢视频列表
//	// 所以如果传入了 token, 此处需要手动调用 JWT 的 mw 校验和解析 token
//	// 但对于 /action/ 是需要校验和解析 token 的
//	reqURI := c.GetRequest().URI().String()
//	if token := c.Query("token"); len(token) != 0 ||
//		strings.Contains(reqURI, "/action/") {
//		jwt.GetInstance().MiddlewareFunc()(ctx, c)
//	}
//})
//// 视频点赞/取消点赞
//_favorite.POST("/action/", FavoriteAction)
//// 喜欢视频列表
//_favorite.GET("/list/", GetFavoriteList)
//// hlog init
//// 配置 hertz 日志格式, 默认格式
//hlog.SetLogger(hlog.DefaultLogger())
//hlog.SetLevel(hlog.LevelInfo)
//// use otel mw
//h.Use(tracing.ServerMiddleware(cfg))
//
//h.Spin()
//}

//func UnsupportedMethod(_ context.Context, c *app.RequestContext) {
//	c.JSON(http.StatusOK, biz.NewFailureResponse("暂不支持该接口服务"))
//}
