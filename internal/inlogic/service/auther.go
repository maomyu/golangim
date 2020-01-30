package service

import (
	"context"
	"time"

	"github.com/yuwe1/golangim/pkg/gerrors"
	"github.com/yuwe1/golangim/pkg/util"
)

type autherService struct {
}

var AuthService = new(autherService)

// 对用户的密钥进行检验
func (*autherService) VerifyToken(ctx context.Context, appId, userId, deviceId int64, token string) error {
	// 根据appid获取app
	app, err := AppService.Get(ctx, appId)
	if err != nil {
		return err
	}
	if app == nil {
		return gerrors.ErrBadRequest
	}
	// 根据私钥和token进行解密
	info, err := util.DecryptToken(token, app.PrivateKey)
	if err != nil {
		return gerrors.ErrBadRequest
	}
	if !(info.AppId == app.Id && info.DeviceId == deviceId && info.UserId == userId) {
		return gerrors.ErrBadRequest
	}
	if info.Expire < time.Now().Unix() {
		return gerrors.ErrBadRequest
	}
	return nil
}
