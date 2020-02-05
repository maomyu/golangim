package dao

import (
	"fmt"

	"github.com/yuwe1/golangim/basic/client/dbpool"
	"github.com/yuwe1/golangim/internal/model"
	"github.com/yuwe1/golangim/pkg/logger"
)

type deviceDao struct {
}

var DeviceDao = new(deviceDao)

func (*deviceDao) UpdateDeviceStatus(deviceId, userId, deviceOnline int64, connAddr string) error {
	s, err, p, c := dbpool.GetSession()
	defer func() {
		if s != nil {
			s.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	s.Begin()
	update := "update device  set user_id = ?,status = ?,conn_addr = ? where device_id = ?"
	_, err = s.Exec(update, userId, deviceOnline, connAddr, deviceId)
	if err != nil {
		s.Rollback()
		return fmt.Errorf("%w", err)
	}
	s.Commit()
	return nil

}

// ListUserOnline 查询用户所有的在线设备
func (*deviceDao) ListOnlineByUserId(appId, userId int64) ([]model.Device, error) {
	s, err, p, c := dbpool.GetSession()
	defer func() {
		if s != nil {
			s.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	s.Begin()
	rows, err := s.Query(
		`select device_id,type,brand,model,system_version,sdk_version,status,conn_addr,create_time,update_time from device where app_id = ? and user_id = ? and status = ?`,
		appId, userId, model.DeviceOnLine)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	devices := make([]model.Device, 0, 5)
	for rows.Next() {
		device := new(model.Device)
		err = rows.Scan(&device.DeviceId, &device.Type, &device.Brand, &device.Model, &device.SystemVersion, &device.SDKVersion,
			&device.Status, &device.ConnAddr, &device.CreateTime, &device.UpdateTime)
		if err != nil {
			logger.Sugar.Error(err)
			return nil, err
		}
		devices = append(devices, *device)
	}
	return devices, nil
}

func (*deviceDao)(d model.Device)error{
	s, err, p, c := dbpool.GetSession()
	defer func() {
		if s != nil {
			s.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	s.Begin()
	insert :="insert into device(device_id,app_id,type,brand,model,system_version,sdk_version,status,conn_addr) values(?,?,?,?,?,?,?,?,?)"
	err = s.Exec(insert,d.DeviceId,
		d.AppId,d.Type,d.Brand,d.Model,d.SystemVersion,
		d.sdk_version,d.Status,""
	)
	if err !=nil{
		return fmt.Errorf("%w",err)
	}
	return nil
}