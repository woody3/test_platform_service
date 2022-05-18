package routers

import (
	"strconv"
	"test_paltform_service/utils"
)

// 所有路由注册完成后，启动gin引擎
func RunGinEngin() {
	registerRouters()
	port := getPort()
	app.StaticFile("favicon.ico", "./favicon.ico")
	app.Run(port)
}

func getPort() string {
	porter := utils.GetConfig().GetInt("server.port")
	port := utils.StringJoin(":", strconv.Itoa(porter))
	return port
}

// 路由组注册
func registerRouters() {
	loadAuthRouters()
}
