package controller

import (
	"net/http"
	"test_platform_service/model"
	"test_platform_service/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Login(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "2001",
			"msg":  "params invalid",
		})
		return
	}
	response := service.UserLoginService(&user)
	ctx.JSON(http.StatusOK, response)
}

func Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "Hello,word"})
}

func Test(ctx *gin.Context) {
	zap.L().Debug("this is Test func")
	ctx.JSON(http.StatusOK, gin.H{"msg": "hahaha"})
}
