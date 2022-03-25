package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"goWeb/service"
	"net/http"
)

type UserLoginController struct {
}

//在使用postman测试时：用户登录和注册可共用一个Controller 目前login等价于register
func (userLogin *UserLoginController) UserLoginController(context *gin.RouterGroup) {
	context.POST("/login", userLogin.UserLoginInstance)
	context.POST("/register", userLogin.UserLoginInstance)
}

func (userLogin *UserLoginController) UserLoginInstance(context *gin.Context) {
	var userLoginService service.UserLoginService
	result, err := userLoginService.UserLoginService(context)
	if err != nil || result == 0 {
		logger.Error("插入数据失败", err)
	} else {
		context.JSON(http.StatusOK, gin.H{
			"message": "注册成功",
		})
	}
}
