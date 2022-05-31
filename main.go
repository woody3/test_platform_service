package main

import (
	"runtime"
	"test_platform_service/routers"
	"test_platform_service/utils"
)

func main() {
	if system := runtime.GOOS; system == "linux" {
		utils.Viper("application-prd.yaml")
	} else {
		utils.Viper("application-dev.yaml")
	}

	routers.EngineBoot(false)
}
