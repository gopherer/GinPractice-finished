package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goWeb/service"
	"net/http"
)

type UserLoginController struct {
}

//在使用postman测试时：用户登录和注册可共用一个Controller 目前login等价于register
func (userLogin *UserLoginController) UserLoginController(context *gin.RouterGroup) {
	context.POST("/login", userLogin.userLoginController)
}

func (userLogin *UserLoginController) userLoginController(context *gin.Context) {
	var userLoginService service.UserLoginService
	err := userLoginService.UserLoginService(context)
	if err == errors.New("输入用户账号不存在") {
		context.JSON(http.StatusOK, gin.H{
			"message": "输入用户账号不存在",
		})
	} else if err == errors.New("用户输入密码有误") {
		context.JSON(http.StatusOK, gin.H{
			"message": "用户输入密码有误",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"message": "用户登陆成功",
		})
	}
}
