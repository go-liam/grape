package job

import (
	"fmt"
	"github.com/go-liam/util/conv"
	"github.com/robfig/cron/v3"
)

func RunHour(c *cron.Cron) {
	println("[INFO]定时器:hour")
	// 每小时的 15分钟
	c.AddFunc("15 * * * *", func() {
		fmt.Println("[EveryHours]1 ", conv.TimeNowFormat(timeFormat))
	})
	// 每小时 的30分钟
	c.AddFunc("30 * * * *", func() {
		fmt.Println("[EveryHours]2 ", conv.TimeNowFormat(timeFormat))
	})
}
