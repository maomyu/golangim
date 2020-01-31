package logic

import "context"

import "github.com/yuwe1/golangim/pkg/pb"

import "github.com/yuwe1/golangim/pkg/contextcli"

import "github.com/yuwe1/golangim/internal/model"

import "github.com/yuwe1/golangim/internal/inlogic/service"

type LogicServerExt struct {
}

func (*LogicServerExt) SendMessage(ctx context.Context, req *pb.SendMessageReq) (*pb.SendMessageResp, error) {
	appId, err := contextcli.GetCtxAppId(ctx)
	if err != nil {
		return nil, err
	}
	// 定义业务服务器的发送
	sender := model.Sender{
		AppId:      appId,
		SenderType: pb.SenderType_ST_BUSINESS,
	}
	return &pb.SendMessageResp{}, service.MessageService.SendMessage(ctx, sender, *req)
}
