package contextcli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/yuwe1/golangim/pkg/logger"
	"google.golang.org/grpc/metadata"
)

const (
	CtxAppId     = "app_id"
	CtxUserId    = "user_id"
	CtxDeviceId  = "device_id"
	CtxToken     = "token"
	CtxRequestId = "request_id"
)

// 存储一个requestId
func ContextWithRequestId(ctx context.Context, requestId int64) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.Pairs(CtxRequestId, strconv.FormatInt(requestId, 10)))
}

// 获得RequestID
func GetContextRequestId(ctx context.Context) int64 {
	// 获取存储结构
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0
	}
	// 检查是否存在requestId
	requestIds, ok := md[CtxRequestId]
	if !ok && len(CtxRequestId) == 0 {
		return 0
	}
	requestId, err := strconv.ParseInt(requestIds[0], 10, 64)
	if err != nil {
		return 0
	}
	return requestId
}

// 获取ctx的用户数据
func GetCtxData(ctx context.Context) (int64, int64, int64, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0, 0, 0, fmt.Errorf("无法识别的用户")
	}
	var (
		appId    int64
		userId   int64
		deviceId int64
		err      error
	)
	appIdStrs, ok := md[CtxAppId]
	if !ok && len(appIdStrs) == 0 {
		return 0, 0, 0, fmt.Errorf("无法识别的用户")
	}
	appId, err = strconv.ParseInt(appIdStrs[0], 10, 64)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, 0, 0, fmt.Errorf("无法识别的用户")
	}
	userIdStrs, ok := md[CtxUserId]
	if !ok && len(userIdStrs) == 0 {
		return 0, 0, 0, fmt.Errorf("无法识别的用户")
	}
	userId, err = strconv.ParseInt(userIdStrs[0], 10, 64)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, 0, 0, fmt.Errorf("无法识别的用户")
	}

	deviceIdStrs, ok := md[CtxDeviceId]
	if !ok && len(deviceIdStrs) == 0 {
		return 0, 0, 0, fmt.Errorf("无法识别的用户")
	}
	deviceId, err = strconv.ParseInt(deviceIdStrs[0], 10, 64)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, 0, 0, fmt.Errorf("无法识别的用户")
	}
	return appId, userId, deviceId, nil
}

// 获取appid
func GetCtxAppId(ctx context.Context) (int64, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0, fmt.Errorf("无法识别的用户")
	}
	appids, ok := md[CtxAppId]
	if !ok && len(appids) == 0 {
		return 0, fmt.Errorf("无法识别的用户")
	}
	appid, err := strconv.ParseInt(appids[0], 10, 64)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, fmt.Errorf("无法识别的用户")
	}
	return appid, nil
}

// 获取token
func GetCtxToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("无法识别的用户")
	}
	tokens, ok := md[CtxToken]
	if !ok && len(tokens) == 0 {
		return "", fmt.Errorf("无法识别的用户")
	}
	return tokens[0], nil
}
