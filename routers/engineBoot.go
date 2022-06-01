package routers

import (
	"fmt"
	"strconv"
	"sync"
	"test_platform_service/middleware"
	"test_platform_service/utils"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine
var _once sync.Once

// 所有路由注册完成后，启动gin引擎
func EngineBoot() {
	genGinInstance()
	registerRouters()
	app.StaticFile("favicon.ico", "./favicon.ico")
	app.Run(getPort())
}

// 生成gin单例
func genGinInstance() {
	conf := utils.GetConfig().Sub("logger")
	err := middleware.InitZapLogger(conf)
	if err != nil {
		panic(fmt.Sprintf("init logger error: %s", err.Error()))
	}
	_once.Do(func() {
		gin.SetMode(conf.GetString("mode"))
		gin.DisableConsoleColor()
		app = gin.New()
		app.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	})
}

func getPort() string {
	porter := utils.GetConfig().GetInt("server.port")
	port := utils.StringJoin(":", strconv.Itoa(porter))
	return port
}
