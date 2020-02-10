package util

import (
	"database/sql"

	"github.com/yuwe1/golangim/pkg/logger"
	"github.com/yuwe1/golangim/pkg/util/sqlid"
)
)

var (
	MessageBodyIdUid *sqlid.Uid
	DeviceIdUid      *sqlid.Uid
)

const (
	DeviceIdBusinessId = "device_id" // 设备id
)

func InitUID(db *sql.DB) {
	var err error
	DeviceIdUid, err = uid.NewUid(db, DeviceIdBusinessId, 5)
	if err != nil {
		logger.Sugar.Error(err)
		panic(err)
	}
}
