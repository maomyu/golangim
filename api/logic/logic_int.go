package logic

import (
	"context"

	"github.com/yuwe1/golangim/pkg/pb"
)

type LogicIntServer struct{}

// SignIn 设备登录
func (*LogicIntServer) SignIn(ctx context.Context, req *pb.SignInReq) (*pb.SignInResp, error) {

	// _, _ = rpc_cli.ConnIntClient.DeliverMessage(contextcli.ContextWithRequestId(context.TODO(), 1), &pb.DeliverMessageReq{
	// 	DeviceId: 12,
	// 	Message: &pb.Message{
	// 		Message: &pb.MessageItem{
	// 			RequestId:      1,
	// 			SenderType:     2,
	// 			SenderId:       1,
	// 			SenderDeviceId: 9,
	// 			ReceiverType:   1,
	// 			ReceiverId:     22,
	// 			MessageBody: &pb.MessageBody{
	// 				MessageType: 1,
	// 				MessageContent: &pb.MessageContent{
	// 					Content: &pb.MessageContent_Text{
	// 						Text: &pb.Text{
	// 							Text: "你好",
	// 						},
	// 					},
	// 				},
	// 			},
	// 			Seq:      1,
	// 			SendTime: time.Now().UnixNano(),
	// 			Status:   1,
	// 		},
	// 	},
	// })
	return &pb.SignInResp{}, nil
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
