package favorite

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	fms "douyin-easy/grpc_gen/favorite"
	clientv3 "go.etcd.io/etcd/client/v3"
	etcdResolver "go.etcd.io/etcd/client/v3/naming/resolver"
)

type ClientConnect struct {
	conn *grpc.ClientConn
}

func NewClientConnect() *ClientConnect {
	fmt.Println("准备连接服务")
	return &ClientConnect{getConnect()}
}
func getConnect() *grpc.ClientConn {

	//注册解析器
	etcdClient, _ := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	etcdResolver, err := etcdResolver.NewBuilder(etcdClient)
	if err != nil {
		panic(err)
	}
	fmt.Println(etcdResolver.Scheme())
	conn, err := grpc.Dial(fmt.Sprintf("etcd:///%s", "douyin/FavoriteServer"),
		grpc.WithResolvers(etcdResolver),
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

func Action(request *fms.FavoriteActionRequest) (*fms.FavoriteActionResponse, error) {
	client := fms.NewFavoriteManagementClient(MyClientConnect.conn)
	reply, err := client.Action(context.Background(), request)
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
func QueryFavorites(request *fms.QueryFavoritesRequest) (*fms.QueryFavoritesResponse, error) {
	client := fms.NewFavoriteManagementClient(MyClientConnect.conn)

	reply, err := client.QueryFavorites(context.Background(), request)
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
