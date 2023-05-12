package main

import (
	"awesomeProject/middleware/db"
	"awesomeProject/routes"
	"database/sql"
	"github.com/go-ini/ini"
	"go.uber.org/zap"
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
	db.InitDb(cfg)
	if db.DB != nil {
		db, _ := db.DB.DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {

			}
		}(db)
	}
	// 加载路由
	r := routes.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
