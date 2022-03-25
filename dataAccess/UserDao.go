package dataAccess

import (
	"github.com/wonderivan/logger"
	"goWeb/model"
	"goWeb/tools"
)

type UserDao struct {
	engine *tools.Orm
}

func (userDao *UserDao) SetUserInfoId(userInfo *model.UserInfo) (int64, error) {
	userDao.engine = tools.DbEngine
	result, err := userDao.engine.InsertOne(userInfo)
	if err != nil {
		logger.Error("userInfo数据添加到数据库失败", err)
		return result, err
	}
	return result, nil
}
func (userDao *UserDao) GetUserInfo(id int64, userInfo *model.UserInfo) (bool, error) {
	userDao.engine = tools.DbEngine
	return userDao.engine.Where("id=?", id).Get(userInfo)
}
func (userDao *UserDao) UpdateUserInfo(userInfo *model.UserInfo) (int64, error) {
	userDao.engine = tools.DbEngine
	result, err := userDao.engine.Where("id=?", userInfo.Id).Update(userInfo)
	if err != nil {
		logger.Error("userInfo数据添加到数据库失败", err)
		return result, err
	}
	return result, nil
}
func (userDao *UserDao) GetUserLoginID(userLogin model.UserLogin) model.UserLogin {
	userDao.engine = tools.DbEngine
	_, err := userDao.engine.Where("user_account", userLogin.UserAccount).Get(userLogin)
	if err != nil {
		logger.Error("GetUserLoginID 失败", err)
	}
	return userLogin
}
func (userDao *UserDao) SetUserLogin(user *model.UserLogin) (int64, error) {
	userDao.engine = tools.DbEngine
	result, err := userDao.engine.InsertOne(user) //result = 0 代表失败， 1 代表成功（存在成功一半的可能）
	if err != nil {
		logger.Error("userLogin 数据添加到数据库失败", err)
		return result, err
	}
	return result, nil
}
