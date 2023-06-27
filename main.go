package main

import (
	"awesomeProject/cron"
	"awesomeProject/middleware/db"
	"awesomeProject/middleware/log"
	"awesomeProject/routes"
	"database/sql"
	"fmt"
	"github.com/go-ini/ini"
	"go.uber.org/zap"
)

func main() {

	// 加载静态配置文件
	cfg, errCfgLoadIni := ini.Load("conf/app.ini")
	cfgSection := cfg.Section("app")
	if errCfgLoadIni != nil {
		log.Logger.Error("fail to load ini",
			zap.Error(errCfgLoadIni))
		panic("fail to load ini")
		return
	}
	// 初始化zap日志
	log.InitLogger(cfg)
	// 加载数据库
	db.InitDb(cfg)
	if db.DB != nil {
		db, _ := db.DB.DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {

			}
		}(db)
	}

	// 定时任务on
	cron.RunCron()
	// 加载路由
	r := routes.SetupRouter()
	// 去除You trusted all proxies, this is NOT safe. 告警
	// https://github.com/gin-gonic/gin/issues/2809
	r.SetTrustedProxies(nil)
	r.Use(log.GinLogger(log.Logger), log.GinRecovery(log.Logger, true))
	// Listen and Server in 0.0.0.0:8080
	err := r.Run(fmt.Sprintf(":%s", cfgSection.Key("httpport")))
	if err != nil {
		return
	}
}
