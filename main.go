package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"goWeb/controller"
	"goWeb/tools"
	"os"
)

func main() {
	cfg, err := tools.ParseConfig("./config/web.json")
	if err != nil {
		logger.Error("tools.ParseConfig 调用失败", err)
		os.Exit(1)
	}
	engine := gin.Default()
	engine.LoadHTMLGlob("Html/user/*")
	_, err = tools.OrmEngine(cfg)
	if err != nil {
		logger.Error("获取数据库实例失败", err)
		os.Exit(1)
	
	}
	userRouter(engine)
	err = engine.Run(cfg.WebHost + ":" + cfg.WebPort)
	if err != nil {
		logger.Crit("engine.Run 调用失败", err)
		os.Exit(1)

	}
}
func userRouter(router *gin.Engine) {
	userGroup := router.Group("/user")
	new(controller.UserRegisterController).UserRegisterController(userGroup)
	new(controller.UserLoginController).UserLoginController(userGroup)
	new(controller.UserInfoController).UserInfoController(userGroup)
}
