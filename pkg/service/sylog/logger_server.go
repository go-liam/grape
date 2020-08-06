package svlog

var Sv InLogger

func init() {
	Sv = new(SvMockLogger)
}

func Init(config string) error {
	return Sv.Init(config)
}

func WriteLog(timeID int64, msg string, level int32) error {
	return Sv.WriteLog(timeID, msg, level)
}

//func Destroy() {
//	 Sv.Destroy()
//}
//
//func Flush() {
//	Sv.Flush()
//}
