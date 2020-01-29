package gerrors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUnknown           = status.New(codes.Unknown, "error unknown").Err() // 服务器未知错误
	ErrUnauthorized      = newError(1002, "error unauthorized")             // 未登录
	ErrNotInGroup        = newError(1003, "error not in group")             // 没有在群组内
	ErrDeviceNotBindUser = newError(1004, "error device not bind user")     // 没有在群组内
	ErrBadRequest        = newError(1005, "error bad request")              // 请求参数错误
	ErrUserAlreadyExist  = newError(1006, "error user already exist")       // 用户已经存在
	ErrGroupAlreadyExist = newError(1007, "error group already exist")      // 群组已经存在
	ErrGroupNotExist     = newError(1008, "error group not exist")          // 群组不存在
	ErrUserNotExist      = newError(1009, "error user not exist")           // 用户不存在
)

func newError(code int64, message string) error {
	return status.New(codes.Code(code), message).Err()
}
