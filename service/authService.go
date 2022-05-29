package service

import (
	"net/http"
	"test_platform_service/model"
	"test_platform_service/repository"
	"test_platform_service/utils"

	"github.com/gin-gonic/gin"
)

func UserLoginService(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "2001",
			"msg":  "params invalid",
		})
		return
	}
	if repository.UserLoginDAO(&user) {
		token, _ := utils.GenToken(user.UserName)
		ctx.JSON(
			http.StatusOK, gin.H{
				"code": "0000",
				"msg":  "success",
				"data": gin.H{"token": token},
			})
		return
	}
	ctx.JSON(
		http.StatusOK, gin.H{
			"code": "1001",
			"msg":  "login failed",
			"data": "",
		})
}
