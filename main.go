package main

import (
	"awesomeProject/routes"
	"fmt"
	"github.com/go-ini/ini"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 初始化zap日志
	logger := zap.NewExample()
	defer logger.Sync()
	// 加载静态配置文件
	cfg, errCfgLoadIni := ini.Load("conf/app.ini")
	if errCfgLoadIni != nil {
		logger.Error("fail to load ini",
			zap.Error(errCfgLoadIni))
		panic("fail to load ini")
		return
	}
	// 加载数据库
	cfgSection := cfg.Section("database")
	dbInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		cfgSection.Key("User"),
		cfgSection.Key("Password"),
		cfgSection.Key("Host"),
		cfgSection.Key("Schema"))
	logger.Info(dbInfo)
	_, dbErr := gorm.Open(mysql.Open(dbInfo), &gorm.Config{})
	if dbErr != nil {
		logger.Error("failed to connect database",
			zap.Error(dbErr))
		panic("failed to connect database")
	}
	// 加载路由
	r := routes.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
