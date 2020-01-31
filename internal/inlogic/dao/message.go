package dao

import (
	"fmt"
	"github.com/yuwe1/golangim/basic/client/dbpool"
	"github.com/yuwe1/golangim/internal/model"
	"github.com/yuwe1/golangim/pkg/logger"
)

type messageDao struct {
}

var MessageDao = new(messageDao)

// 添加一条消息
func (*messageDao) Add(tablename string, message model.Message) error {
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
	insert := fmt.Sprintf(`insert into %s(app_id,object_type,object_id,request_id,sender_type,sender_id,sender_device_id,receiver_type,receiver_id,to_user_ids,type,content,seq,send_time) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)`, tablename)
	_, err = s.Exec(insert, message.AppId, message.ObjectType, message.ObjectId, message.RequestId, message.SenderType, message.SenderId, message.SenderDeviceId,
		message.ReceiverType, message.ReceiverId, message.ToUserIds, message.Type, message.Content, message.Seq, message.SendTime)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
