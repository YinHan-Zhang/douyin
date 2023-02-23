package configs

import "time"

/*
 @Author: 71made
 @Date: 2023/02/14 22:43
 @ProductName: etcd.go
 @Description:
*/
const TCP = "tcp"
const Etcd = "etcd"
const HTTP = "http"
const EtcdPort = ":2379"

// const EtcdURL = ServerAddrPrefix + EtcdPort
const EtcdURL = ServerAddrPrefix + EtcdPort
const EtcdDialTimeout = 5 * time.Second
const EtcdTTL = 10
const ServerIP = "127.0.0.1"
const (
	ServerNamePrefix = "douyin"
	ServerAddrPrefix = HTTP + "://" + ServerIP
)
const UserServer = "UserServer"
const VideoServer = "VideoServer"
const (
	UserServerName  = ServerNamePrefix + "/" + UserServer
	VideoServerName = ServerNamePrefix + "/" + VideoServer
)
