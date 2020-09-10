package job

import (
	"fmt"
	"github.com/go-liam/util/conv"
	"github.com/robfig/cron/v3"
)

func RunDay(c *cron.Cron) {
	println("[INFO]定时器:day")
	// 每天 3点 20分
	c.AddFunc("20 3 * * *", func() {
		fmt.Println("[EveryDay]2", conv.TimeNowFormat(timeFormat))
		//dayStepWX(conv.TimesTampToday(),  conv.TimesTampTomorrow())
	})
	// 每天 4点 30分
	c.AddFunc("30 4 * * *", func() {
		fmt.Println("[EveryDay]3", conv.TimeNowFormat(timeFormat))
		//dayStepAct(conv.TimesTampToday(),  conv.TimesTampTomorrow())
	})
	// 每天 5点 40分
	c.AddFunc("40 5 * * *", func() {
		fmt.Println("[EveryDay]E", conv.TimeNowFormat(timeFormat))
	})
}
