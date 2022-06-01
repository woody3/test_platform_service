package main

import (
	"runtime"
	"test_platform_service/routers"
	"test_platform_service/utils"
	"test_platform_service/repository"
)

func main() {
	if system := runtime.GOOS; system == "linux" {
		utils.Viper("application-prd.yaml")
	} else {
		utils.Viper("application-dev.yaml")
	}
//         初始化数据库
//         repository.InitDb()
	routers.EngineBoot()
}
