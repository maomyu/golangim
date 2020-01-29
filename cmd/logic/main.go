package main

import (
	"github.com/yuwe1/golangim/api/logic"
	"github.com/yuwe1/golangim/basic"
	"github.com/yuwe1/golangim/basic/config"
	"github.com/yuwe1/golangim/pkg/logger"
	"github.com/yuwe1/golangim/pkg/rpc_cli"
)

func main() {
	basic.Init()

	rpc_cli.InitConnIntClient(config.GetLogicConfig().GetConnRpcAddrs())
	logic.StartRpcServer()
	logger.Logger.Info("logic server start")
	select {}
}
