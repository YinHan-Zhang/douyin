package favorite

import (
	fms "douyin-easy/grpc_gen/favorite"
	"os"
	"testing"
)

var reqAction *fms.FavoriteActionRequest
var reqQueryFavorites *fms.QueryFavoritesRequest

var resAction *fms.FavoriteActionResponse
var resQueryFavorite *fms.QueryFavoritesResponse

func TestMain(m *testing.M) {
	//测试前：数据装载、配置初始化等前置工作
	//构造Action请求数据
	reqAction = &fms.FavoriteActionRequest{
		Type:    fms.Action_Cancel,
		UserId:  1,
		VideoId: 2,
	}
	//构造Action响应数据
	resAction = &fms.FavoriteActionResponse{}
	//构造QueryFavorites请求数据
	reqQueryFavorites = &fms.QueryFavoritesRequest{
		UserId:   1,
		VideoIds: nil,
	}
	//favoriteList := []*fms.Favorite{
	//	{
	//		Id:         1,
	//		UserId:     1,
	//		VideoId:    1,
	//		IsFavorite: true,
	//	},
	//}
	//构造QueryFavorites响应数据
	resQueryFavorite = &fms.QueryFavoritesResponse{
		FavoriteList: nil,
	}
	code := m.Run()

	//测试后：释放资源等收尾工作
	//...
	defer MyClientConnect.conn.Close()
	os.Exit(code)
}
func TestAction(t *testing.T) {
	reply, err := Action(reqAction)
	if err != nil {
		t.Error(err)
	}
	t.Log(reply)

}
func TestQueryFavorites(t *testing.T) {
	reply, err := QueryFavorites(reqQueryFavorites)
	if err != nil {
		t.Error(err)
	}
	t.Log(reply)
}
