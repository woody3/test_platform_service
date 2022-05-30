package repository

import (
	"test_platform_service/model"
	"test_platform_service/utils"
)

func UserLoginDAO(user *model.User) (flag bool) {
	var count int64
	db := GetGormDbInstance()
	db.Where("username = ? and password= ?", user.UserName, utils.EncryptMD5(user.Password)).Find(user).Count(&count)
	if count == 1 {
		flag = true
	}
	return
}
