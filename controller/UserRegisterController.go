package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"goWeb/service"
	"net/http"
)

type UserRegisterController struct {
}

func (userRegister *UserRegisterController) UserRegisterController(context *gin.RouterGroup) {
	context.POST("/register", userRegister.postUserRegister)
	context.GET("register", userRegister.getUserRegister)
}

func (userRegister *UserRegisterController) postUserRegister(context *gin.Context) {
	var userRegisterService service.UserRegisterService
	result, err := userRegisterService.UserRegisterService(context)
	if err != nil || result == 0 {
		logger.Error("插入数据失败", err)
		context.JSON(http.StatusOK, gin.H{
			"message": "用户账号已被使用",
		})
	} else {
		context.HTML(http.StatusOK, "user-register-success.html", gin.H{})
	
}
func (userRegister *UserRegisterController) getUserRegister(context *gin.Context) {
	context.HTML(http.StatusOK, "user-register.html", gin.H{})
}
