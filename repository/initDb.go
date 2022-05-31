package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"test_platform_service/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gormDb *gorm.DB
var sqlDb *sql.DB
var _once sync.Once

func InitDb() {
	var err error
	var config = utils.GetConfig().Sub("dataSource")
	_once.Do(func() {
		var dsn = utils.StringJoin(config.GetString("username"), ":", config.GetString("password"), "@tcp(", config.GetString("url"), ")/", config.GetString("database.0"), "?charset=utf8&parseTime=True&loc=Local")
		loggers := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{})
		gormDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: loggers})
		if err != nil {
			panic(fmt.Sprintf("connect DataBase error: %s", err.Error()))
		}
		sqlDb, _ := gormDb.DB()
		sqlDb.SetMaxOpenConns(config.GetInt("conns.maxOpenConns"))
		sqlDb.SetMaxIdleConns(config.GetInt("conns.maxIdleConns"))
	})
}
