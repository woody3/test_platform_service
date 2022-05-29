package controller

import (
	"net/http"
	"test_platform_service/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Login(ctx *gin.Context) {
	service.UserLoginService(ctx)
	ctx.JSON(http.StatusOK, gin.H{"msg": "login success"})
}

func Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "Hello,word"})
}

func Test(ctx *gin.Context) {
	zap.L().Debug("this is Test func")
	ctx.JSON(http.StatusOK, gin.H{"msg": "hahaha"})
}
