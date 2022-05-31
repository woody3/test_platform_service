package service

import (
	"test_platform_service/model"
	"test_platform_service/repository"
	"test_platform_service/utils"

	"github.com/gin-gonic/gin"
)

func UserLoginService(user *model.User) *gin.H {
	if repository.UserLoginDAO(user) {
		token, _ := utils.GenToken(user.UserName)
		return &gin.H{
			"code": "0000",
			"msg":  "success",
			"data": gin.H{"token": token},
		}
	}
	return &gin.H{
		"code": "1001",
		"msg":  "user not exists",
		"data": "",
	}
}
