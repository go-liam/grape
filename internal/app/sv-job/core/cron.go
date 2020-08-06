package core

import (
	"fmt"
	"github.com/go-liam/util/conv"
	"github.com/robfig/cron/v3"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

// 冥操作，可以重复操作

/*
https://godoc.org/github.com/robfig/cron
Field name   | Mandatory? | Allowed values  | Allowed special characters
----------   | ---------- | --------------  | --------------------------
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?
*/
func JobHour() {
	// 执行(Minutes,Hours,Day of month,Month,Day of week )
	println("[cron]定时器", conv.TimeNowFormat(timeFormat))
	nyc, _ := time.LoadLocation("Asia/Chongqing")
	c := cron.New(cron.WithLocation(nyc))
	//c := cron.New()
	// 每1分钟
	c.AddFunc("*/1 * * * *", func() {
		fmt.Println("[EveryMinutes]1 ", conv.TimeNowFormat(timeFormat))
	})
	//每10分钟
	c.AddFunc("*/10 * * * *", func() {
		fmt.Println("[EveryMinutes]2 ", conv.TimeNowFormat(timeFormat))
	})
	// 每小时的 15分钟
	c.AddFunc("15 * * * *", func() {
		fmt.Println("[EveryHours]1 ", conv.TimeNowFormat(timeFormat))
	})
	// 每小时 的30分钟
	c.AddFunc("30 * * * *", func() {
		fmt.Println("[EveryHours]2 ", conv.TimeNowFormat(timeFormat))
	})
	// 每天 2点 10分
	c.AddFunc("10 2 * * *", func() {
		fmt.Println("[EveryDay]1", conv.TimeNowFormat(timeFormat))
		//dayStepAPI(conv.TimesTampToday(),  conv.TimesTampTomorrow())
	})
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
	go c.Start()
}
