package user

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	usrv "douyin-easy/grpc_gen/user"
	"douyin-easy/pkg/configs"

	clientv3 "go.etcd.io/etcd/client/v3"
	etcdResolver "go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
@Author: 71made
@Date: 2023/02/19 13:41
@ProductName: client.go
@Description: 构建 gRPC 客户端，用于测试（可以不通过 etcd 服务发现机制获取到服务地址，通过最简单的方式构建连接即可）
*/

type ClientConnect struct {
	conn *grpc.ClientConn
}

func NewClientConnect() *ClientConnect {
	return &ClientConnect{getConnect()}
}

func getConnect() *grpc.ClientConn {

	//注册解析器
	etcdClient, _ := clientv3.New(clientv3.Config{
		Endpoints:   []string{configs.EtcdURL},
		DialTimeout: configs.EtcdDialTimeout,
	})
	etcd_resolver, err := etcdResolver.NewBuilder(etcdClient)
	if err != nil {
		panic(err)
	}
	fmt.Println(etcd_resolver.Scheme())
	conn, err := grpc.Dial(fmt.Sprintf("etcd:///%s", configs.UserServerName),
		grpc.WithResolvers(etcd_resolver),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		log.Fatal("服务端出错，连接不上：", err)
		return nil

	}
	return conn

}

var MyClientConnect = NewClientConnect()

func CheckLoginUser(request *usrv.CheckLoginUserRequest) (*usrv.CheckLoginUserResponse, error) {
	client := usrv.NewUserManagementClient(MyClientConnect.conn)
	reply, err := client.CheckLoginUser(context.Background(), request)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	buf, err := json.MarshalIndent(reply, "", "\t")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println(string(buf))
	return reply, nil
}

func CreateUser(request *usrv.CreateUserRequest) (*usrv.CreateUserResponse, error) {
	client := usrv.NewUserManagementClient(MyClientConnect.conn)

	reply, err := client.CreateUser(context.Background(), request)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	buf, err := json.MarshalIndent(reply, "", "\t")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println(string(buf))
	return reply, nil
}

func QueryUsers(request *usrv.QueryUsersRequest) (*usrv.QueryUsersResponse, error) {
	client := usrv.NewUserManagementClient(MyClientConnect.conn)

	reply, err := client.QueryUsers(context.Background(), request)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	buf, err := json.MarshalIndent(reply, "", "\t")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println(string(buf))
	return reply, nil
}

func QueryUser(request *usrv.QueryUserRequest) (*usrv.QueryUserResponse, error) {
	client := usrv.NewUserManagementClient(MyClientConnect.conn)

	reply, err := client.QueryUser(context.Background(), request)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	buf, err := json.MarshalIndent(reply, "", "\t")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println(string(buf))
	return reply, nil
}
