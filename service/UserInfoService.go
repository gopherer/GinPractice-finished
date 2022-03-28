package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"goWeb/dataAccess"
	"goWeb/model"
	"goWeb/tools"
	"net/http"
)

type UserInfoService struct {
}

func (userInfoService *UserInfoService) UserInfoServicePost(context *gin.Context) (int64, error) {
	var userInfo model.UserInfo
	var userRegister = GetUserRegisterService()
	userInfoDao := dataAccess.UserInfoDao{}
	userRegisterDao := dataAccess.UserRegisterDao{}
	userRegisterDao.GetUserRegisterID(&userRegister)
	userInfo.Id = userRegister.Id
	//err := context.BindJSON(&userInfo)
	err := context.Bind(&userInfo)
	if err != nil {
		logger.Error("BindJSON 用户绑定失败", err)
		context.JSON(http.StatusOK, gin.H{
			"message": "输入数据有误请重新输入",
		})
		return 0, errors.New("输入数据有误请重新输入")
	}
	result := tools.UserJsonLen(userInfo)
	if !result {
		context.JSON(http.StatusOK, gin.H{
			"message": "昵称或个人介绍过长",
		})
	}
	return userInfoDao.UpdateUserInfo(&userInfo)
}

func (userInfoService *UserInfoService) UserInfoServiceGet(id int64, userInfo *model.UserInfo) (bool, error) {
	var userDao dataAccess.UserInfoDao
	return userDao.GetUserInfo(id, userInfo)
}
