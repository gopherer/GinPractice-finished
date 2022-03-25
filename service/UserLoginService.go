package service

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"goWeb/dataAccess"
	"goWeb/model"
	"goWeb/tools"
)

type UserLoginService struct {
}

var userLogin model.UserLogin

func (userLoginService *UserLoginService) UserLoginService(context *gin.Context) (int64, error) {
	userLogin.Id = tools.GetUserId()
	err := context.BindJSON(&userLogin)
	if err != nil {
		logger.Error("userLogin BindJson 失败", err)
		return 0, err
	}
	//将用户注册数据插入数据库表user_login同时会在user_info表中将新建一条数据把id置为用户注册时的id号,以此关联二者
	userInfo := model.UserInfo{
		Id:       userLogin.Id,
		UserName: "",
		UserBio:  "",
	}
	userDao := dataAccess.UserDao{}
	_, _ = userDao.SetUserInfoId(&userInfo)
	return userDao.SetUserLogin(&userLogin)
}
func GetUserLoginService() model.UserLogin {
	return userLogin
}
