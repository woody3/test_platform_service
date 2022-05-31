package repository

import (
	"test_platform_service/model"
	"test_platform_service/utils"
)

func UserLoginDAO(user *model.User) (flag bool) {
	var count int64
	gormDb.Table("user_info").Select("user_info.id").Where("username = ? and password= ?", user.UserName, utils.EncryptMD5(user.Password)).Count(&count)
	if count == 1 {
		flag = true
	}
	return
}
