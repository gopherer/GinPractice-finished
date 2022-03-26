package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"goWeb/model"
	"goWeb/service"
	"net/http"
)

type UserInfoController struct {
}

func (userInfoController *UserInfoController) UserInfoController(context *gin.RouterGroup) {
	context.POST("/info", whetherActivateUserInfoMiddleWare(), userInfoController.userInfoPost)
	context.GET("/info", whetherActivateUserInfoMiddleWare(), userInfoController.userInfoGet)
}

func (userInfoController *UserInfoController) userInfoPost(context *gin.Context) {
	var userInfoService service.UserInfoService
	result, err := userInfoService.UserInfoServicePost(context)
	if err != nil || result == 0 {
		logger.Error("插入数据失败", err)
	} else {
		context.JSON(http.StatusOK, gin.H{
			"message": "userInfo数据插入成功",
		})
	}
}
func (userInfoController *UserInfoController) userInfoGet(context *gin.Context) {
	var userInfo model.UserInfo
	var userInfoService service.UserInfoService
	userRegister := service.GetUserRegisterService()
	result, err := userInfoService.UserInfoServiceGet(userRegister.Id, &userInfo)
	if err != nil || result == false {
		logger.Error("获取数据失败", err)
	} else {
		context.JSON(http.StatusOK, &userInfo)
	}
}

//中间件
func whetherActivateUserInfoMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		userRegister := service.GetUserRegisterService()
		//多写一点，严谨一点
		if userRegister.Id == 0 && userRegister.UserAccount == "" && userRegister.UserPassWord == "" {
			context.JSON(http.StatusOK, gin.H{
				"message": "请先登入才可以访问哦",
			})
		}
	}
}
