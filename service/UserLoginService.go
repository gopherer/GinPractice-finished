package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"goWeb/dataAccess"
	"goWeb/model"
)

type UserLoginService struct {
}

func (userLoginService *UserLoginService) UserLoginService(context *gin.Context) error {
	var userLogin model.UserLogin
	err := context.BindJSON(&userLogin)
	if err != nil {
		logger.Error("用户数据格式不匹配", err)
		return err
	}
	userLoginCopy := userLogin
	userLoginDao := dataAccess.UserLoginDao{}
	_, err = userLoginDao.GetUserLogin(&userLoginCopy)
	if err != nil {
		logger.Error("查询数据库表失败", err)
	}
	if userLogin.UserAccount != userLoginCopy.UserAccount {
		return errors.New("输入用户账号不存在")
	} else if userLogin.UserAccount == userLoginCopy.UserAccount && userLogin.UserPassWord != userLoginCopy.UserPassWord {
		return errors.New("用户输入密码有误")
	}
	//将userLogin变为userRegister  就可以对userInfo操作
	userRegister.Id = userLoginCopy.Id
	userRegister.UserAccount = userLoginCopy.UserAccount
	userRegister.UserPassWord = userLoginCopy.UserPassWord
	return nil
}
