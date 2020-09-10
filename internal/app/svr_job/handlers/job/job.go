package job

import (
	"github.com/go-liam/util/conv"
	"github.com/robfig/cron/v3"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

func RunJob() {
	// 执行(Minutes,Hours,Day of month,Month,Day of week )
	println("[cron]定时器", conv.TimeNowFormat(timeFormat))
	nyc, _ := time.LoadLocation("Asia/Chongqing")
	c := cron.New(cron.WithLocation(nyc))
	// 分钟
	RunMinutes(c)
	// hour
	RunHour(c)
	// day
	RunDay(c)
	// run
	go c.Start()
}
