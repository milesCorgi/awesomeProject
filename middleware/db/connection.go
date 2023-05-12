package db

import (
	"fmt"
	"github.com/go-ini/ini"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDb(cfg *ini.File) {
	logger := zap.NewExample()
	defer logger.Sync()
	cfgSection := cfg.Section("database")
	dbInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		cfgSection.Key("User"),
		cfgSection.Key("Password"),
		cfgSection.Key("Host"),
		cfgSection.Key("Schema"))
	logger.Info(dbInfo)
	db, dbErr := gorm.Open(mysql.Open(dbInfo), &gorm.Config{})
	if dbErr != nil {
		logger.Error("failed to connect database",
			zap.Error(dbErr))
		panic("failed to connect database")
	}
	sqlDB, _ := db.DB()
	if err := sqlDB.Ping(); err != nil {
		logger.Error(err.Error())
		panic("failed to ping database")
	}
	DB = db
}
