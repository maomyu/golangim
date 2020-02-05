package main

import (
	"github.com/yuwe1/golangim/api/wsconn"
	"github.com/yuwe1/golangim/basic"
	"github.com/yuwe1/golangim/basic/config"
	"github.com/yuwe1/golangim/internal/ws_conn"
	"github.com/yuwe1/golangim/pkg/rpc_cli"
	"github.com/yuwe1/golangim/pkg/util"
)

func main() {
	basic.Init()

	go func() {
		defer util.RecoverPanic()
		wsconn.StartRPCServer()
	}()
	rpc_cli.InitlogicIntClient(config.GetWsconfConfig().GetLogicRPCAddrs())
	ws_conn.StartWSServer(config.GetWsconfConfig().GetWsListenAddr())
}
