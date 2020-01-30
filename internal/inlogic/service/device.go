package service

import (
	"context"

	"github.com/yuwe1/golangim/internal/inlogic/dao"
)

// 设备
type deviceService struct {
}

const (
	deviceOnline = 1
	deviceoffie  = 0
)

var DeviceService = new(deviceService)

func (*deviceService) Online(ctx context.Context, appId, deviceId, userId int64, connectAddr string) error {
	err := dao.DeviceDao.UpdateDeviceStatus(deviceId, userId, deviceOnline, connectAddr)
	if err != nil {
		return err
	}

	// 设置缓存
	return nil
}
