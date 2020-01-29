package ws_conn

import (
	"context"

	"github.com/yuwe1/golangim/pkg/contextcli"
	"github.com/yuwe1/golangim/pkg/logger"
	"github.com/yuwe1/golangim/pkg/pb"
)

func DeliverMessage(ctx context.Context, req *pb.DeliverMessageReq) error {
	// 获取设备对应的TCP连接
	conn := load(req.DeviceId)
	if conn == nil {
		logger.Sugar.Warn("ctx id nil")
		return nil
	}

	// 发送消息
	conn.OutPut(pb.PackageType_PT_MESSAGE, contextcli.GetContextRequestId(ctx), nil, req.Message)
	return nil
}
