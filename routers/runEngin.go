package routers

// 路由组注册
func registerRouters() {
	addAuthRouters()
}

// 所有路由注册完成后，启动gin引擎
func RunGinEngin() {
	registerRouters()
	eng.Run()
}
