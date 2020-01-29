package main

import "github.com/yuwe1/golangim/pkg/util"

import "github.com/yuwe1/golangim/api/wsconn"

import "github.com/yuwe1/golangim/pkg/rpc_cli"

import "github.com/yuwe1/golangim/basic/config"

import "github.com/yuwe1/golangim/internal/ws_conn"

import "github.com/yuwe1/golangim/basic"

func main() {
	basic.Init()
	go func() {
		defer util.RecoverPanic()
		wsconn.StartRPCServer()
	}()
	rpc_cli.InitlogicIntClient(config.GetWsconfConfig().GetLogicRPCAddrs())
	// rpc_cli.InitConnIntClient(config.GetLogicConfig().GetConnRpcAddrs())
	ws_conn.StartWSServer(config.GetWsconfConfig().GetWsListenAddr())
}
