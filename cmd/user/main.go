package main

import (
	"douyin-easy/cmd/user/command"
	"fmt"
	"log"
	"net"
	"os"

	db "douyin-easy/cmd/user/dal"

	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"
)

/*
 @Author: 71made
 @Date: 2023/02/19 13:35
 @ProductName: main.go
 @Description: 服务启动函数
*/

var (
	ServiceName  string
	ServiceAddr  string
	EtcdAddress  string
	Argon2Config *command.Argon2Params
)

// 初始化config文件
func InitConfig() *viper.Viper {
	workDir, _ := os.Getwd()
	var v = viper.New()
	// 读取的文件名
	v.SetConfigName("userConfig")
	// 读取的文件类型
	v.SetConfigType("yml")
	// 读取的路径
	fmt.Println(workDir)
	v.AddConfigPath(workDir + "\\config")

	err := v.ReadInConfig()
	if err != nil {
		// fmt.Println(workDir)
		fmt.Println(err)
		panic("Read dbConfig faild!")
	}
	return v
}

// User RPC Server 端配置初始化
func Init(v *viper.Viper) {
	db.Init()
	Argon2Config = &command.Argon2Params{
		Memory:      v.GetUint32("Server.Argon2ID.Memory"),
		Iterations:  v.GetUint32("Server.Argon2ID.Iterations"),
		Parallelism: uint8(v.GetUint("Server.Argon2ID.Parallelism")),
		SaltLength:  v.GetUint32("Server.Argon2ID.SaltLength"),
		KeyLength:   v.GetUint32("Server.Argon2ID.KeyLength"),
	}
	ServiceName = v.GetString("Server.Name")
	ServiceAddr = fmt.Sprintf("%s:%d", v.GetString("Server.Address"), v.GetInt("Server.Port"))
	EtcdAddress = fmt.Sprintf("%s:%d", v.GetString("Etcd.Address"), v.GetInt("Etcd.Port"))
}

// User RPC Server 端运行
func main() {
	var v = InitConfig()
	Init(v)
	svr := newServer()
	grpclog.Info("Running user grpc server...")
	//开启监听
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("server started...")
	err = svr.Serve(listen)
	if err != nil {
		grpclog.Fatal("User grpc server start failed: ", err)
		return
	}

}
