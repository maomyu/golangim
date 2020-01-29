package basic

import (
	"github.com/yuwe1/golangim/basic/client/dbpool"
	"github.com/yuwe1/golangim/basic/client/rediscli/redispool"
	"github.com/yuwe1/golangim/basic/config"
	"github.com/yuwe1/golangim/basic/mq"
)

func Init() {
	config.Init()
	// rediscli.Init()
	redispool.Init()
	mq.Init()
	dbpool.Init()
}
