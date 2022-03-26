package service

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"goWeb/dataAccess"
	"goWeb/model"
	"goWeb/tools"
	"strconv"
)

type UserRegisterService struct {
}

var userRegister model.UserRegister

func (userRegisterService *UserRegisterService) UserRegisterService(context *gin.Context) (int64, error) {
	userRegister.Id = tools.GetUserId()
	err := context.BindJSON(&userRegister)
	if err != nil {
		logger.Error("userRegister BindJson 失败", err)
		return 0, err
	}
	//将用户注册数据插入数据库表user_Register同时会在user_info表中将新建一条数据把id置为用户注册时的id号,以此关联二者
	//UserName 需要设置默认值 以确保主键不会冲突
	userInfo := model.UserInfo{
		Id:       userRegister.Id,
		UserName: strconv.Itoa(int(userRegister.Id)),
		UserBio:  "",
	}
	userInfoDao := dataAccess.UserInfoDao{}
	userRegisterDao := dataAccess.UserRegisterDao{}
	_, _ = userInfoDao.SetUserInfoIdName(&userInfo)
	return userRegisterDao.SetUserRegister(&userRegister)
}
func GetUserRegisterService() model.UserRegister {
	return userRegister
}
