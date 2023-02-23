package pack

import (
	"douyin-easy/cmd/favorite/dal"
	fsvr "douyin-easy/grpc_gen/favorite"
)

func Favorite(f *dal.Favorite) *fsvr.Favorite {
	if f == nil {
		return nil
	}

	return &fsvr.Favorite{
		Id:         int64(f.ID),
		UserId:     int64(f.UserId),
		VideoId:    int64(f.VideoId),
		IsFavorite: f.IsFavorite(),
	}
}

func Favorites(fs []*dal.Favorite) []*fsvr.Favorite {
	if fs == nil || len(fs) == 0 {
		return make([]*fsvr.Favorite, 0)
	}

	res := make([]*fsvr.Favorite, len(fs))
	for i, f := range fs {
		res[i] = Favorite(f)
	}

	return res
}
