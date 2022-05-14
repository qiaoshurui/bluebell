package mysql

import "errors"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNoExist     = errors.New("用户不存在")
	ErrorInvalidPassWord = errors.New("账号或密码错误")
	ErrorInvalidID       = errors.New("无效的ID")
)
