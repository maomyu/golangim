package rpc_cli

import (
	"context"

	"github.com/yuwe1/golangim/pkg/logger"
	"github.com/yuwe1/golangim/pkg/pb"
	"google.golang.org/grpc"
)

// 定义rpc客户端
var (
	LogicIntClient pb.LogicIntClient
	ConnIntClient  pb.ConnIntClient
)

// 初始化客户端【拨号】
func InitlogicIntClient(addr string) {
	conn, err := grpc.DialContext(context.TODO(), addr, grpc.WithInsecure())
	if err != nil {
		logger.Sugar.Error(err)
		panic(err)
	}
	LogicIntClient = pb.NewLogicIntClient(conn)
}

// 初始化连接层的客户端
func InitConnIntClient(addr string) {
	conn, err := grpc.DialContext(context.TODO(), addr, grpc.WithInsecure())
	if err != nil {
		logger.Sugar.Error(err)
		panic(err)
	}
	ConnIntClient = pb.NewConnIntClient(conn)
}
