package svlog

/*
https://en.wikipedia.org/wiki/Syslog
值	严重程度	关键词	不推荐使用的关键字	描述	健康）状况
0	紧急情况	emerg	panic[7]	系统无法使用	恐慌的状况。[8]
1	警报	alert		必须立即采取行动	应立即纠正的状况，例如损坏的系统数据库。[8]
2	危急	crit		临界条件	硬设备错误。[8]
3	错误	err	error[7]	错误条件
4	警告	warning	warn[7]	警告条件
5	注意	notice		正常但重要的条件	不是错误条件，但可能需要特殊处理的条件。[8]
6	信息性	info		信息性消息
7	除错	debug		调试级消息	包含通常仅在调试程序时使用的信息的消息。[8]
*/

// RFC5424 log message levels.
const (
	LevelEmergency     = 0 // iota
	LevelAlert         = 1
	LevelCritical      = 2 // 危急
	LevelError         = 3
	LevelWarning       = 4 // 警告
	LevelNotice        = 5
	LevelInformational = 6
	LevelDebug         = 7
)

type InLogger interface {
	Init(config string) error
	WriteLog(logID int64, msg string, level int32) error
	//Destroy()
	//Flush()
}
