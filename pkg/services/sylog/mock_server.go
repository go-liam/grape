package svlog

import "log"

type SvMockLogger struct {
}

func (sv *SvMockLogger) Init(config string) error {
	return nil
}

func (sv *SvMockLogger) WriteLog(logID int64, msg string, level int32) error {
	log.Printf("[%d]-%d-%s", level, logID, msg)
	return nil
}

func (sv *SvMockLogger) Destroy() {
	println("Destroy")
}

func (sv *SvMockLogger) Flush() {
	println("Flush")
}
