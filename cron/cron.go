package cron

import (
	"awesomeProject/middleware/caches"
	"awesomeProject/middleware/log"
	"github.com/robfig/cron/v3"
	"time"
)

func RunCron() {
	c := cron.New()

	_, err := c.AddFunc("@every 1s", func() {
		cache := caches.GetAllCache()
		for k, v := range cache {
			if v.ExpiredTime.Before(time.Now()) {
				caches.DelACache(k)
				log.Logger.Info("del key:" + k + "ExpiredTime: " + v.ExpiredTime.Format("2006/01/02 15:04"))
			}
		}
	})
	if err != nil {
		panic("fail AddFunc")
	}

	c.Start()
	time.Sleep(time.Second * 5)
}
