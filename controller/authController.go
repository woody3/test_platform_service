package controller

import (
	"net/http"
	"test_paltform_service/service"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var msg string = service.UserLoginService()
	c.JSON(http.StatusOK, gin.H{"msg": msg})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Hello,word"})
}

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "hahaha"})
}
