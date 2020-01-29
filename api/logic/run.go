package logic

import (
	"net"

	"github.com/yuwe1/golangim/basic/config"
	"github.com/yuwe1/golangim/pkg/pb"
	"github.com/yuwe1/golangim/pkg/util"
	"google.golang.org/grpc"
)

// StartRpcServer 启动rpc服务
func StartRpcServer() {
	go func() {
		defer util.RecoverPanic()

		intListen, err := net.Listen("tcp", config.GetLogicConfig().GetRpcIntListenAddr())
		if err != nil {
			panic(err)
		}
		intServer := grpc.NewServer()
		pb.RegisterLogicIntServer(intServer, &LogicIntServer{})
		err = intServer.Serve(intListen)
		if err != nil {
			panic(err)
		}
	}()

}
