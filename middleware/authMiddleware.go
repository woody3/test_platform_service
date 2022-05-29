package middleware

import (
	"net/http"
	"strings"
	"test_platform_service/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			ctx.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(http.StatusOK, gin.H{
				"code": "1002",
				"msg":  "请求头中auth格式有误",
			})
			ctx.Abort()
			return
		}

		mc, err := utils.ParseToken(parts[1])
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			ctx.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文ctx上
		// 后续的处理函数可以用过ctx.Get("username")来获取当前请求的用户信息
		ctx.Set("username", mc.Username)
		ctx.Next()
	}
}
