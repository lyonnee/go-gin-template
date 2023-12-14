package logic

import "errors"

var (
	ERR_INVALID_PASSWORD   = errors.New("无效的密码")
	ERR_LOGIC_EXECUTE_FAIL = errors.New("逻辑执行错误")
	ERR_INVALID_HASH       = errors.New("无效的hash密文")
)
