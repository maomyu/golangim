package service

import (
	"context"

	"github.com/yuwe1/golangim/pkg/pb"
)

type logicService struct{}

var LogicService = new(logicService)

// 设备登录
func (l *logicService) SignIn(ctx context.Context, req *pb.SignInReq) error {
	// 验证token
	err := AuthService.VerifyToken(ctx, req.AppId, req.UserId, req.DeviceId, req.Token)
	if err != nil {
		return err
	}

	// 标记用户上线(即设备上线)
	err = DeviceService.Online(ctx, req.AppId, req.DeviceId, req.UserId, req.ConnAddr)
	if err != nil {
		return err
	}

	return nil
}
