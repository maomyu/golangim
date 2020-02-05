package main

import (
	"fmt"
	"time"

	"github.com/yuwe1/golangim/basic"
	"github.com/yuwe1/golangim/basic/client/dbpool"
	"github.com/yuwe1/golangim/pkg/logger"
	"github.com/yuwe1/golangim/pkg/util/sqlid"
)

// 注册uid
func NewUid(businessid string, len int) *sqlid.Uid {
	s, err, p, c := dbpool.GetSession()
	defer func() {
		if p != nil {
			s.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	fmt.Println(s.ID)
	u, _ := sqlid.NewUid(s.DB, businessid, len)
	return u
}
func main() {
	basic.Init()
	u := NewUid("test", 10)
	for {
		fmt.Println(u.Get())
		time.Sleep(1 * time.Second)
	}
}
