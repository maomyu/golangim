package service

import "context"

import "github.com/yuwe1/golangim/internal/inlogic/cache"

type seqService struct {
}

var SeqService = new(seqService)

// 从缓存中获取下一个序列号
func (*seqService) GetUserNextSeq(ctx context.Context, appId, userId int64) (int64, error) {
	return cache.SeqCache.Incr(cache.SeqCache.UserKey(appId, userId))
}
