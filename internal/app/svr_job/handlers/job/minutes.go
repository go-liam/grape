package job

import (
	"fmt"
	"github.com/go-liam/util/conv"
	"github.com/robfig/cron/v3"
	"grape/internal/app/svr_job/handlers/site"
)

func RunMinutes(c *cron.Cron) {
	println("[INFO]定时器:minutes")
	// 每1分钟
	c.AddFunc("*/1 * * * *", func() {
		go fmt.Println("[EveryMinutes]1 ", conv.TimeNowFormat(timeFormat))
	})
	// 每5分钟
	c.AddFunc("*/1 * * * *", func() {
		go site.GetAllSiteHealth()
	})
	//每10分钟
	c.AddFunc("*/10 * * * *", func() {
		go fmt.Println("[EveryMinutes]2 ", conv.TimeNowFormat(timeFormat))
	})
}
