package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"goWeb/service"
	"net/http"
)

type UserRegisterController struct {
}

//在使用postman测试时：用户登录和注册可共用一个Controller 目前Register等价于register
func (userRegister *UserRegisterController) UserRegisterController(context *gin.RouterGroup) {
	context.POST("/register", userRegister.UserRegisterInstance)
}

func (userRegister *UserRegisterController) UserRegisterInstance(context *gin.Context) {
	var userRegisterService service.UserRegisterService
	result, err := userRegisterService.UserRegisterService(context)
	if err != nil || result == 0 {
		logger.Error("插入数据失败", err)
	} else {
		context.JSON(http.StatusOK, gin.H{
			"message": "注册成功",
		})
	}
}
