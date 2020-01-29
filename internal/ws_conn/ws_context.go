package ws_conn

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/micro/protobuf/proto"
	"github.com/yuwe1/golangim/pkg/logger"
	"github.com/yuwe1/golangim/pkg/pb"
	"github.com/yuwe1/golangim/pkg/util"
	"google.golang.org/grpc/status"
)

// websocket客户端
type WSConnContext struct {
	Conn     *websocket.Conn
	AppId    int64
	DeviceId int64
	UserId   int64
}

// 新建一个websocket客户端
func NewWSConnContext(conn *websocket.Conn, appId, userId, deviceId int64) *WSConnContext {
	return &WSConnContext{
		Conn:     conn,
		AppId:    appId,
		UserId:   userId,
		DeviceId: deviceId,
	}
}

// 处理连接
func (c *WSConnContext) DoConn() {
	defer util.RecoverPanic()

	for {
		err := c.Conn.SetReadDeadline(time.Now().Add(12 * time.Minute))
		_, data, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		// c.Conn.WriteJSON("hahah" + string(data))
		c.HandlePackage(data)
	}
}

// HandlePackage 处理请求发包
func (c *WSConnContext) HandlePackage(bytes []byte) {
	var input pb.Input
	err := proto.Unmarshal(bytes, &input)
	if err != nil {
		logger.Sugar.Error(err)
		c.Release()
	}

	switch input.Type {
	case pb.PackageType_PT_SYNC:
		// c.Sync(input)
	case pb.PackageType_PT_HEARTBEAT:
		// c.Heartbeat(input)
	case pb.PackageType_PT_MESSAGE:
		// c.MessageACK(input)
	default:
		logger.Logger.Info("switch other")
	}

}

// 投递消息
func (c *WSConnContext) OutPut(pt pb.PackageType, requestId int64, err error, message proto.Message) {
	var output = pb.Output{
		Type:      pt,
		RequestId: requestId,
	}
	if err != nil {
		stat, _ := status.FromError(err)
		output.Code = int32(stat.Code())
		output.Message = stat.Message()
	}
	if message != nil {
		msgbytes, err := proto.Marshal(message)
		if err != nil {
			logger.Sugar.Error(err)
			return
		}
		output.Data = msgbytes
	}
	outputbytes, err := proto.Marshal(&output)
	if err != nil {
		logger.Sugar.Error(err)
		return
	}
	// c.Conn.WriteJSON("我是机器人" + string(outputbytes))
	err = c.Conn.WriteMessage(websocket.BinaryMessage, outputbytes)
	if err != nil {
		logger.Sugar.Error(err)
		return
	}
}

// Release 释放TCP连接
func (c *WSConnContext) Release() {
	// 从本地manager中删除tcp连接

	// 关闭tcp连接

	// 通知业务服务器设备下线

}
