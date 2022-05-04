package repository

import (
	"fmt"
	"log"
	"os"
	"test_paltform_service/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDb() {
	var err error
	var config = utils.GetConfig()
	var dsn = utils.StringJoin(config.GetString("dataBase.username"), ":", config.GetString("dataBase.password"), "@tcp(", config.GetString("dataBase.url"), ")/", config.GetString("dataBase.dbName.0"), "?charset=utf8&parseTime=True&loc=Local")
	loggers := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{})
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: loggers})
	if err != nil {
		panic(fmt.Sprintf("connect DataBase error: %s", err.Error()))
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxOpenConns(40)
	sqlDb.SetMaxIdleConns(20)
}

func GetDb() *gorm.DB {
	return db
}
