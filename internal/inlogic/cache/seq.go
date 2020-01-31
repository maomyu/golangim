package cache

import "strconv"

import "github.com/yuwe1/golangim/basic/client/rediscli/redispool"

import "github.com/yuwe1/golangim/pkg/logger"

import "github.com/garyburd/redigo/redis"

import "fmt"

type seqCache struct {
}

var SeqCache = new(seqCache)

const (
	UserSeqKey  = "user_seq"
	GroupSeqKey = "group_seq"
)

func (*seqCache) UserKey(appId, userId int64) string {
	return UserSeqKey + strconv.FormatInt(appId, 10) + ":" + strconv.FormatInt(userId, 10)
}

func (*seqCache) GroupKey(appId, groupId int64) string {
	return GroupSeqKey + strconv.FormatInt(appId, 10) + ":" + strconv.FormatInt(groupId, 10)
}

// 将序列号加1
func (*seqCache) Incr(key string) (int64, error) {
	p, err, r, c := redispool.NewSession()
	defer func() {
		if p != nil {
			p.Relase(r, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	conn := p.GetConn()
	seq, err := redis.Int64(conn.Do("INCR", key))
	if err != nil {
		return 0, fmt.Errorf("incr:[%w]", err)
	}
	return seq, nil
}
