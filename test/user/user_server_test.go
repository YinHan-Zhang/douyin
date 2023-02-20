package user

import (
	usrv "douyin-easy/grpc_gen/user"
	"errors"
	"os"
	"testing"
)

/*
 @Author: 71made
 @Date: 2023/02/19 13:42
 @ProductName: user_server_test.go
 @Description: 通过调用 client 中函数获取客户端，并调用测试 server 服务功能
*/

var reqCheckLoginUser *usrv.CheckLoginUserRequest
var reqCreateUser *usrv.CreateUserRequest
var reqQueryUsers *usrv.QueryUsersRequest
var reqQueryUser *usrv.QueryUserRequest

var resCheckLoginUser *usrv.CheckLoginUserResponse
var resCreateUser *usrv.CreateUserResponse
var resQueryUsers *usrv.QueryUsersResponse
var resQueryUser *usrv.QueryUserResponse

func TestMain(m *testing.M) {
	//测试前：数据装载、配置初始化等前置工作
	//构造CheckLoginUser请求数据
	reqCheckLoginUser = &usrv.CheckLoginUserRequest{
		Username: "xiaoni",
		Password: "demo520",
	}
	//构造CheckLoginUser响应数据
	resCheckLoginUser = &usrv.CheckLoginUserResponse{
		UserId: 1,
	}

	//构造CreateUser请求
	reqCreateUser = &usrv.CreateUserRequest{
		Username: "didi",
		Password: "zky",
		Avatar:   "static/avatar/empty_avatar.jpeg",
	}
	//构造CreateUserr响应数据
	resCreateUser = &usrv.CreateUserResponse{
		User: &usrv.User{
			Id:   4,
			Name: "didi",
		},
	}

	//构造QueryUsers请求
	reqQueryUsers = &usrv.QueryUsersRequest{
		UserIds: []int64{1, 2, 3, 4},
	}
	//构造QueryUsers响应数据
	userList := []*usrv.User{
		{
			Id:   1,
			Name: "xiaoni",
		},
		{
			Id:   2,
			Name: "daimao",
		},
		{
			Id:   3,
			Name: "cyc",
		},
		{
			Id:   4,
			Name: "didi",
		},
	}
	resQueryUsers = &usrv.QueryUsersResponse{
		UserList: userList,
	}

	//构造QueryUser请求
	reqQueryUser = &usrv.QueryUserRequest{
		UserId: 1,
	}
	//构造QueryUser响应数据
	resQueryUser = &usrv.QueryUserResponse{
		User: &usrv.User{
			Name: "xiaoni",
		},
	}

	code := m.Run()

	//测试后：释放资源等收尾工作
	//...
	defer MyClientConnect.conn.Close()
	os.Exit(code)
}

func TestCheckLoginUser(t *testing.T) {
	reply, err := CheckLoginUser(reqCheckLoginUser)
	if err != nil {
		t.Error(err)
	}
	if reply.UserId == resCheckLoginUser.UserId {
		t.Log("CheckLoginUser test success")
	} else {
		t.Error(errors.New("CheckLoginUser test faild"))
	}
}

func TestCreateUser(t *testing.T) {
	reply, err := CreateUser(reqCreateUser)
	if err != nil {
		t.Error(err)
	}
	if reply.User.Id == resCreateUser.User.Id && reply.User.Name == resCreateUser.User.Name {
		t.Log("CreateUser test success")
	} else {
		t.Error(errors.New("CreateUser test faild"))
	}
}
func TestQueryUsers(t *testing.T) {
	reply, err := QueryUsers(reqQueryUsers)
	if err != nil {
		t.Error(err)
	}
	count := 0
	for index, user := range resQueryUsers.UserList {
		if user.Id != reply.UserList[index].Id || user.Name != reply.UserList[index].Name {
			t.Error(errors.New("QueryUsers test faild"))
		} else {
			count += 1
		}
	}
	if count == 4 {
		t.Log("QueryUsers test success")
	}
}
func TestQueryUser(t *testing.T) {
	reply, err := QueryUser(reqQueryUser)
	if err != nil {
		t.Error(err)
	}
	if reply.User.Name == resQueryUser.User.Name {
		t.Log("QueryUser test success")
	} else {
		t.Error(errors.New("QueryUser test faild"))
	}
}
