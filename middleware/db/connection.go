package db

import (
	"awesomeProject/middleware/log"
	"fmt"
	"github.com/go-ini/ini"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/sharding"
)

var (
	DB *gorm.DB
)

func InitDb(cfg *ini.File) {

	cfgSection := cfg.Section("database")
	dbInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true",
		cfgSection.Key("User"),
		cfgSection.Key("Password"),
		cfgSection.Key("Host"),
		cfgSection.Key("Schema"))
	log.Logger.Info(dbInfo)
	db, dbErr := gorm.Open(mysql.Open(dbInfo), &gorm.Config{})
	db.Use(sharding.Register(sharding.Config{
		ShardingKey:         "source_id",
		NumberOfShards:      64,
		PrimaryKeyGenerator: sharding.PKSnowflake,
	}, "video_lists"))
	if dbErr != nil {
		log.Logger.Error("failed to connect database",
			zap.Error(dbErr))
		panic("failed to connect database")
	}
	sqlDB, _ := db.DB()
	if err := sqlDB.Ping(); err != nil {
		log.Logger.Error(err.Error())
		panic("failed to ping database")
	}
	DB = db
}
