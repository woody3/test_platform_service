package routers

import (
	"sync"
	"test_paltform_service/controller"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine
var _once sync.Once

func init() {
	_once.Do(func() {
		app = gin.Default()
	})
}

func addAuthRouters() {
	authGroup := app.Group("auth")
	{
		authGroup.POST("/login", controller.Login)
		authGroup.POST("/logout", controller.Logout)
		authGroup.GET("/test", controller.Test)
	}

}
