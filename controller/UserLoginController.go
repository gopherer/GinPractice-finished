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
	errExist := errors.New("输入用户账号不存在")
	errUnExist := errors.New("用户输入密码有误")
	//即便文本相同，每次调用errors.New()函数都会返回一个不同的错误值。因此不能直接拿error类型做比较
	if err == nil {
		context.HTML(http.StatusOK, "user-login-success.html", gin.H{})
	} else if err.Error() == errExist.Error() {
		context.JSON(http.StatusOK, gin.H{
			"message": "输入用户账号不存在",
		})
	} else if err.Error() == errUnExist.Error() {
		context.JSON(http.StatusOK, gin.H{
			"message": "用户输入密码有误",
		})
	}
}
