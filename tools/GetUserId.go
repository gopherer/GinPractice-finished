package tools

import (
	"github.com/wonderivan/logger"
)

var Id int64

func GetUserId() int64 {
	engine := DbEngine
	_, err := engine.SQL("select count(*) from user_login").Get(&Id)
	if err != nil {
		logger.Error("从数据库获取userID 失败", err)
	}
	Id++
	return Id
}
