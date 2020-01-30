package dao

import (
	"fmt"
	"github.com/yuwe1/golangim/basic/client/dbpool"
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
