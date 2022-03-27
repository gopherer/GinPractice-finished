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

type UserLoginService struct {
}

func (userLoginService *UserLoginService) UserLoginService(context *gin.Context) error {
	var userLogin model.UserLogin
	err := context.BindJSON(&userLogin)
	if err != nil {
		logger.Error("用户数据格式不匹配", err)
		return err
	}
	result := tools.UserJsonLen(userLogin)
	if !result {
		context.JSON(http.StatusOK, gin.H{
			"message": "账号密码至少大于6位,不得大于20位",
		})
		return errors.New("账号密码至少大于6位,不得大于20位")
	}
	userLoginDao := dataAccess.UserLoginDao{}
	result = userLoginDao.GetUserLogin(&userLogin)
	if !result {
		logger.Error("查询数据库表失败", err)
	}
	if userLogin.UserAccount == "" && userLogin.Id == 0 {
		return errors.New("输入用户账号不存在")
	} else if userLogin.UserAccount != "" && userLogin.UserPassWord == "" && userLogin.Id != 0 {
		return errors.New("用户输入密码有误")
	}
	//将userLogin变为userRegister  就可以对userInfo操作
	userRegister.Id = userLogin.Id
	userRegister.UserAccount = userLogin.UserAccount
	userRegister.UserPassWord = userLogin.UserPassWord
	return nil
}
