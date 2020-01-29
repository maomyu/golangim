package wsconn

import (
	"context"
	"fmt"
	"net"

	"github.com/yuwe1/golangim/basic/config"
	"github.com/yuwe1/golangim/internal/ws_conn"
	"github.com/yuwe1/golangim/pkg/pb"
	"google.golang.org/grpc"
)

type ConnIntServer struct {
}

// 投递消息
func (s *ConnIntServer) DeliverMessage(ctx context.Context, req *pb.DeliverMessageReq) (*pb.DeliverMessageResp, error) {
	fmt.Println("开始投递消息")
	return &pb.DeliverMessageResp{}, ws_conn.DeliverMessage(ctx, req)
}

// 启动rpc服务
func StartRPCServer() {
	listener, err := net.Listen("tcp", config.GetWsconfConfig().GetRpcListenAddr())
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterConnIntServer(server, &ConnIntServer{})
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
