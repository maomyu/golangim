package service

import (
	"context"

	"github.com/yuwe1/golangim/internal/inlogic/dao"
	"github.com/yuwe1/golangim/internal/model"

	"github.com/yuwe1/golangim/pkg/contextcli"
	"github.com/yuwe1/golangim/pkg/logger"
	"github.com/yuwe1/golangim/pkg/pb"
	"github.com/yuwe1/golangim/pkg/rpc_cli"
	"github.com/yuwe1/golangim/pkg/util"
	"go.uber.org/zap"
)

type messageService struct {
}

var MessageService = new(messageService)

func (*messageService) SendMessage(ctx context.Context, sender model.Sender, req *pb.SendMessageReq) error {
	switch req.ReceiverType {
	case pb.ReceiverType_RT_USER:
		// 发送者的类型为user
		if sender.SenderType == pb.SenderType_ST_USER {

		} else {
			err := MessageService.SendToUser(ctx, sender, req.ReceiverId, 0, req)
			if err != nil {
				return err
			}
		}
		// / 否则为其他
	}
	return nil
}

func (*messageService) SendToUser(ctx context.Context,
	sender model.Sender,
	toUserId int64,
	roomSeq int64, req *pb.SendMessageReq) error {
	logger.Logger.Debug("message_store_send_to_usr",
		zap.String("message_id", req.MessageId),
		zap.Int64("app_id", sender.AppId),
		zap.Int64("to_user_id", toUserId),
	)
	var (
		seq = roomSeq
		err error
	)
	if req.IsPersist {
		// 获取user的下一个序列号{ppId:userId}
		seq, err = SeqService.GetUserNextSeq(ctx, sender.AppId, toUserId)
		if err != nil {
			return err
		}
		//  分解出messageType和messageContent
		messageType, messageContent := model.PBToMessageBody(req.MessageBody)

		selfMessage := model.Message{
			AppId: sender.AppId,
			// 发送的目标类型
			ObjectType:     model.MessageObjectTypeUser,
			ObjectId:       toUserId,
			RequestId:      contextcli.GetContextRequestId(ctx),
			SenderType:     int32(sender.SenderType),
			SenderId:       sender.SenderId,
			SenderDeviceId: sender.DeviceId,
			ReceiverType:   int32(req.ReceiverType),
			ReceiverId:     req.ReceiverId,
			ToUserIds:      model.FormatUserIds(req.ToUserIds),
			Type:           messageType,
			Content:        messageContent,
			Seq:            seq,
			SendTime:       util.UnunixMilliTime(req.SendTime),
			Status:         int32(pb.MessageStatus_MS_NORMAL),
		}
		// / 将该消息添加到数据库中
		err = MessageService.Add(ctx, selfMessage)
		if err != nil {
			return err
		}
	}
	messageItem := pb.MessageItem{
		RequestId:      contextcli.GetContextRequestId(ctx),
		SenderType:     sender.SenderType,
		SenderId:       sender.SenderId,
		SenderDeviceId: sender.DeviceId,
		ReceiverType:   req.ReceiverType,
		ReceiverId:     req.ReceiverId,
		ToUserIds:      req.ToUserIds,
		MessageBody:    req.MessageBody,
		Seq:            seq,
		SendTime:       req.SendTime,
		Status:         pb.MessageStatus_MS_NORMAL,
	}

	// 查询用户在线设备
	devices, err := DeviceService.ListOnlineBUserId(ctx, sender.AppId, toUserId)
	if err != nil {
		return err
	}

	for i := range devices {
		// 消息不需要投递给发送消息的设备
		if sender.DeviceId == devices[i].DeviceId {
			continue
		}
		err = MessageService.sendToDevice(ctx, devices[i], messageItem)
		if err != nil {
			return err
		}
	}
	return nil
}

// / SendToDevice 将消息发送给设备
func (*messageService) sendToDevice(ctx context.Context, device model.Device, msgItem pb.MessageItem) error {
	if device.Status == model.DeviceOnLine {
		message := pb.Message{Message: &msgItem}
		_, err := rpc_cli.ConnIntClient.DeliverMessage(contextcli.ContextWithAddr(ctx, device.ConnAddr), &pb.DeliverMessageReq{
			DeviceId: device.DeviceId, Message: &message})
		if err != nil {
			return err
		}
	}

	// todo 其他推送厂商
	return nil
}

// / Add 添加消息
func (*messageService) Add(ctx context.Context, message model.Message) error {
	return dao.MessageDao.Add("message", message)
}

// / ListOnlineByUserId 获取用户的所有在线设备
func (*deviceService) ListOnlineBUserId(ctx context.Context, appId, userId int64) ([]model.Device, error) {
	// 从缓存中获取
	devices, err := dao.DeviceDao.ListOnlineByUserId(appId, userId)
	if err != nil {
		return nil, err
	}

	// 设置缓存
	return devices, nil
}
