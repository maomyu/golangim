package logic

import (
	"context"

	"github.com/yuwe1/golangim/internal/inlogic/service"
	"github.com/yuwe1/golangim/pkg/logger"
	"github.com/yuwe1/golangim/pkg/pb"
)

type LogicIntServer struct{}

// SignIn 设备登录

func (*LogicIntServer) SignIn(ctx context.Context, req *pb.SignInReq) (*pb.SignInResp, error) {
	logger.Sugar.Info(req.UserId, "的设备", req.DeviceId, "登录设备")
	return &pb.SignInResp{}, service.LogicService.SignIn(ctx, req)
}

// Sync 设备同步消息
func (*LogicIntServer) Sync(ctx context.Context, req *pb.SyncReq) (*pb.SyncResp, error) {
	// messages, err := service.MessageService.ListByUserIdAndSeq(ctx, req.AppId, req.UserId, req.Seq)
	// if err != nil {
	// 	return nil, err
	// }
	return &pb.SyncResp{}, nil
}

// MessageACK 设备收到消息ack
func (*LogicIntServer) MessageACK(ctx context.Context, req *pb.MessageACKReq) (*pb.MessageACKResp, error) {
	return &pb.MessageACKResp{}, nil
}

// Offline 设备离线
func (*LogicIntServer) Offline(ctx context.Context, req *pb.OfflineReq) (*pb.OfflineResp, error) {
	return &pb.OfflineResp{}, nil
}
