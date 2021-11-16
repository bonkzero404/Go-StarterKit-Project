package database

import (
	"fmt"
	"go-boilerplate-clean-arch/config"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	var err error
	var dbType string

	if config.Config("DB_DRIVER") == "mysql" {
		DB, err = gorm.Open(mysql.Open(DsnMySqlDB()))
	} else if config.Config("DB_DRIVER") == "pgsql" {
		DB, err = gorm.Open(postgres.Open(DsnPostgreSqlDB()))
	} else {
		panic("Database driver not available")
	}

	dbType = config.Config("DB_DRIVER")

	if err != nil {
		errMessage := fmt.Sprintf("Failed to connect database %s", dbType)
		panic(errMessage)
	}

	dbPooling(DB)

	return DB
}

func dbPooling(sqlDb *gorm.DB) error {
	sqlDB, err := DB.DB()

	if err != nil {
		panic("failed to connect database")
	}

	maxIdleConsConf := config.Config("DB_MAX_IDLE_CONNS")
	maxOpenConsConf := config.Config("DB_MAX_OPEN_CONNS")

	maxIdleCons, _ := strconv.Atoi(maxIdleConsConf)
	maxOpenCons, _ := strconv.Atoi(maxOpenConsConf)

	sqlDB.SetMaxIdleConns(maxIdleCons)

	sqlDB.SetMaxOpenConns(maxOpenCons)

	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}
