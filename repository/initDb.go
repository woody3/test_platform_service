package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"test_paltform_service/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gormDb *gorm.DB
var _once sync.Once

func initDb() {
	var err error
	var config = utils.GetConfig()
	var dsn = utils.StringJoin(config.GetString("dataSource.username"), ":", config.GetString("dataSource.password"), "@tcp(", config.GetString("dataSource.url"), ")/", config.GetString("dataSource.database.0"), "?charset=utf8&parseTime=True&loc=Local")
	loggers := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{})
	gormDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: loggers})
	if err != nil {
		panic(fmt.Sprintf("connect DataBase error: %s", err.Error()))
	}
	sqlDb, _ := gormDb.DB()
	sqlDb.SetMaxOpenConns(config.GetInt("maxOpenConns"))
	sqlDb.SetMaxIdleConns(config.GetInt("maxIdleConns"))
}

func GetGormDbInstance() *gorm.DB {
	_once.Do(func() {
		initDb()
	})
	return gormDb
}

func GetSqlDbInstance() *sql.DB {
	sqldb, err := GetGormDbInstance().DB()
	if err != nil {
		panic(fmt.Sprintf("Get sqldb error: %s", err.Error()))
	}
	return sqldb
}
