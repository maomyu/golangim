package service

import (
	"context"

	"github.com/yuwe1/golangim/internal/inlogic/dao"
	"github.com/yuwe1/golangim/internal/model"
	"github.com/yuwe1/golangim/pkg/gerrors"
	"github.com/yuwe1/golangim/pkg/logger"
	"github.com/yuwe1/golangim/pkg/util"
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

// 注册设备
func (*deviceService) Register(ctx context.Context, device model.Device) (int, error) {
	// 根据appid获取设备的信息
	app, err := AppService.Get(ctx, device.AppId)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, err
	}
	if app == nil {
		return 0, gerrors.ErrBadRequest
	}
	// 获取deviceId
	deviceId, err := util.DeviceIdUid.Get()
	if err != nil {
		return 0, err
	}
	device.DeviceId = deviceId
	// 在数据库中添加一个数据库
	err = dao.DeviceDao.Add(device)
	return 0, gerrors.ErrBadRequest
}
