package dataAccess

import (
	"github.com/wonderivan/logger"
	"goWeb/model"
	"goWeb/tools"
)

type UserInfoDao struct {
	engine *tools.Orm
}

func (userInfoDao *UserInfoDao) SetUserInfoIdName(userInfo *model.UserInfo) (int64, error) {
	userInfoDao.engine = tools.DbEngine
	result, err := userInfoDao.engine.InsertOne(userInfo)
	if err != nil {
		logger.Error("userInfo数据添加到数据库失败", err)
		return result, err
	}
	return result, nil
}
func (userInfoDao *UserInfoDao) GetUserInfo(id int64, userInfo *model.UserInfo) (bool, error) {
	userInfoDao.engine = tools.DbEngine
	return userInfoDao.engine.Where("id=?", id).Get(userInfo)
}
func (userInfoDao *UserInfoDao) UpdateUserInfo(userInfo *model.UserInfo) (int64, error) {
	userInfoDao.engine = tools.DbEngine
	has, _ := userInfoDao.engine.Table("user_info").Where("user_name = ?", userInfo.UserName).Exist()
	if has {
		return -1, nil
	}
	result, err := userInfoDao.engine.Where("id = ?", userInfo.Id).Update(userInfo)
	if err != nil {
		logger.Error("userInfo数据添加到数据库失败", err)
		return result, err
	}
	return result, nil
}
