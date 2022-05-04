package main

import (
	"runtime"
	"test_paltform_service/repository"
	"test_paltform_service/utils"
)

func main() {
	if system := runtime.GOOS; system == "linux" {
		utils.Viper("application-prd.yaml")
	} else {
		utils.Viper("application-dev.yaml")
	}
	repository.InitDb()
}
