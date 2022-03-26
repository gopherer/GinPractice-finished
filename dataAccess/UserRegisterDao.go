package dataAccess

import (
	"github.com/wonderivan/logger"
	"goWeb/model"
	"goWeb/tools"
)

type UserRegisterDao struct {
	engine *tools.Orm
}

func (userRegisterDao *UserRegisterDao) GetUserRegisterID(userRegister *model.UserRegister) {
	userRegisterDao.engine = tools.DbEngine
	_, err := userRegisterDao.engine.Where("user_account", userRegister.UserAccount).Get(userRegister)
	if err != nil {
		logger.Error("GetUserRegisterID 失败", err)
	}
}
func (userRegisterDao *UserRegisterDao) SetUserRegister(user *model.UserRegister) (int64, error) {
	userRegisterDao.engine = tools.DbEngine
	result, err := userRegisterDao.engine.InsertOne(user) //result = 0 代表失败， 1 代表成功（存在成功一半的可能）
	if err != nil {
		logger.Error("userRegister 数据添加到数据库失败", err)
		return result, err
	}
	return result, nil
}
