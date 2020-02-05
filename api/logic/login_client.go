package logic

import (
	"context"

	"github.com/yuwe1/golangim/internal/model"
	"github.com/yuwe1/golangim/pkg/gerrors"
	"github.com/yuwe1/golangim/pkg/pb"
)

// 定义一个结构体
type LogicClient struct {
}

// 注册一个设备
func (*LogicClient) RegisterDevice(ctx context.Context, req *pb.RegisterDeviceReq) (*pb.RegisterDeviceResp, error) {
	device := model.Device{
		AppId:         req.AppId,
		Type:          req.Type,
		Brand:         req.Brand,
		Model:         req.Model,
		SystemVersion: req.SystemVersion,
		SDKVersion:    req.SdkVersion,
	}
	// 检查参数
	if device.AppId == 0 || device.Type == 0 || device.Brand == "" || device.Model == "" ||
		device.SystemVersion == "" || device.SDKVersion == "" {
		return nil, gerrors.ErrBadRequest
	}
	// 注册设备,需要返回deviceid

	return &pb.RegisterDeviceResp{}, nil
}

// 添加一个用户
func (*LogicClient) AddUser(ctx context.Context, req *pb.AddUserReq) (*pb.AddUserResp, error) {
	return &pb.AddUserResp{}, nil
}

// 获取一个用户的用户信息
func (*LogicClient) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserResp, error) {
	return &pb.GetUserResp{}, nil
}

// 发送消息
func (*LogicClient) SendMessage(ctx context.Context, req *pb.SendMessageReq) (*pb.SendMessageResp, error) {
	return &pb.SendMessageResp{}, nil
}
