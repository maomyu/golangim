package ws_conn

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/yuwe1/golangim/basic/config"
	"github.com/yuwe1/golangim/pkg/contextcli"
	"github.com/yuwe1/golangim/pkg/gerrors"
	"github.com/yuwe1/golangim/pkg/logger"
	"github.com/yuwe1/golangim/pkg/pb"
	"github.com/yuwe1/golangim/pkg/rpc_cli"
	"google.golang.org/grpc/status"
)

// 定义升级
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 65536,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// handler函数
func ws(w http.ResponseWriter, r *http.Request) {
	// appId, _ := strconv.ParseInt(r.Header.Get(contextcli.CtxAppId), 10, 64)
	// userId, _ := strconv.ParseInt(r.Header.Get(contextcli.CtxUserId), 10, 64)
	// deviceId, _ := strconv.ParseInt(r.Header.Get(contextcli.CtxDeviceId), 10, 64)
	// token := r.Header.Get(contextcli.CtxToken)
	// requestId, _ := strconv.ParseInt(contextcli.CtxRequestId, 10, 64)
	appId, _ := strconv.ParseInt(r.FormValue(contextcli.CtxAppId), 10, 64)
	userId, _ := strconv.ParseInt(r.FormValue(contextcli.CtxUserId), 10, 64)
	deviceId, _ := strconv.ParseInt(r.FormValue(contextcli.CtxDeviceId), 10, 64)
	token := r.FormValue(contextcli.CtxToken)
	requestId, _ := strconv.ParseInt(contextcli.CtxRequestId, 10, 64)
	requestId = 1
	// 对参数进行检查
	if appId == 0 || userId == 0 ||
		deviceId == 0 || token == "" || requestId == 0 {
		s, _ := status.FromError(gerrors.ErrUnauthorized)
		bytes, err := json.Marshal(s.Proto)
		if err != nil {
			logger.Sugar.Error(err)
		}
		w.Write(bytes)
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Sugar.Error(err)
		return
	}

	// 断开之前的连接
	preCtx := load(deviceId)
	if preCtx != nil {
		preCtx.DeviceId = -1
	}
	ctx := NewWSConnContext(conn, appId, deviceId, userId)
	store(deviceId, ctx)

	// 调用logic  rpc
	_, err = rpc_cli.LogicIntClient.SignIn(contextcli.ContextWithRequestId(context.TODO(), requestId), &pb.SignInReq{
		AppId:    appId,
		UserId:   userId,
		DeviceId: deviceId,
		Token:    token,
		ConnAddr: config.GetWsconfConfig().GetLocalAddr(),
	})
	if err != nil {
		s, _ := status.FromError(gerrors.ErrUnauthorized)
		bytes, err := json.Marshal(s.Proto)
		if err != nil {
			logger.Sugar.Error(err)
		}
		w.Write(bytes)
	}
	// test 投递消息

	ctx.DoConn()

}
