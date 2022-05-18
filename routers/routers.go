package routers

import (
	"io"
	"os"
	"sync"
	"test_paltform_service/controller"
	"test_paltform_service/utils"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine
var _once sync.Once

func init() {
	_once.Do(func() {
		gin.DisableConsoleColor()
		f, _ := os.Create(utils.GetConfig().GetString("logDir"))
		gin.DefaultWriter = io.MultiWriter(f)
		app = gin.Default()
	})
}

// 配置路由
func loadAuthRouters() {
	authGroup := app.Group("auth")
	{
		authGroup.POST("/login", controller.Login)
		authGroup.POST("/logout", controller.Logout)
		authGroup.GET("/test", controller.Test)
	}

}
