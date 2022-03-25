package service

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"goWeb/dataAccess"
	"goWeb/model"
	"net/http"
)

type UserInfoService struct {
}

func (userInfoService *UserInfoService) UserInfoServicePost(context *gin.Context) (int64, error) {
	var userInfo model.UserInfo
	var userLogin = GetUserLoginService()
	userDao := dataAccess.UserDao{}
	userLogin = userDao.GetUserLoginID(userLogin)
	userInfo.Id = userLogin.Id
	err := context.BindJSON(&userInfo)
	if err != nil {
		logger.Error("BindJSON 用户绑定失败", err)
		context.JSON(http.StatusOK, gin.H{
			"message": "输入数据有误请重新输入",
		})
		return 0, err
	}
	return userDao.UpdateUserInfo(&userInfo)
}

func (userInfoService *UserInfoService) UserInfoServiceGet(id int64, userInfo *model.UserInfo) (bool, error) {
	var userDao dataAccess.UserDao
	return userDao.GetUserInfo(id, userInfo)
}
