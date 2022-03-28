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
	context.GET("/set-info", whetherActivateUserInfoMiddleWare(), userInfoController.redirectSetUserInfo)

}

func (userInfoController *UserInfoController) userInfoPost(context *gin.Context) {
	var userInfoService service.UserInfoService
	result, err := userInfoService.UserInfoServicePost(context)
	if result == 0 {
		logger.Error("插入数据失败", err)
	} else if result == -1 {
		context.JSON(http.StatusOK, gin.H{
			"message": "用户名已存在",
		})
	} else {
		context.HTML(http.StatusOK, "set-user-info-after-get.html", gin.H{})
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
		context.HTML(http.StatusOK, "get-user-info.html", gin.H{
			"userName": userInfo.UserName,
			"userBio":  userInfo.UserBio,
		})
	}
}

func (userInfoController *UserInfoController) redirectSetUserInfo(context *gin.Context) {
	context.HTML(http.StatusOK, "set-user-info.html", gin.H{})
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
			//abort（）顾名思义就是终止的意思，也就是说执行该函数，会终止后面所有的该请求下的函数。
			context.Abort()
		}
	}
}
