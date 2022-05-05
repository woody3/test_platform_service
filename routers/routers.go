package routers

import (
	"sync"
	"test_paltform_service/controller"

	"github.com/gin-gonic/gin"
)

var eng *gin.Engine
var lock sync.Mutex

func singleEng() {
	lock.Lock()
	defer lock.Unlock()
	eng = gin.Default()
}

func addAuthRouters() {
	singleEng()
	authGroup := eng.Group("auth")
	{
		authGroup.POST("/login", controller.Login)
		authGroup.POST("/logout", controller.Logout)
		authGroup.POST("/test", controller.Test)
	}

}
