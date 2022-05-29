package routers

import (
	"test_platform_service/controller"
	"test_platform_service/middleware"
)

// 配置路由
func loadAuthRouters() {
	authGroup := app.Group("auth")
	{
		authGroup.POST("/login", controller.Login)
		authGroup.POST("/logout", middleware.AuthMiddleware(), controller.Logout)
		authGroup.GET("/test", controller.Test)
	}
}
