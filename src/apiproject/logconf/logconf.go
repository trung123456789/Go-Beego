package logconf

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"path"
	"runtime"
	"time"
)

type Logger struct {
	Logger     *logs.BeeLogger
	Context    string
	Identifier int64
}

func NewLogger(ctx string) *Logger {
	identifier := time.Now().UnixNano()
	return &Logger{
		Logger:     logs.GetBeeLogger(),
		Context:    ctx,
		Identifier: identifier,
	}
}

type delegate func(format string, v ...interface{})

func (log *Logger) fmtLog(fn delegate, format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	_, filename := path.Split(file)
	fmtString := fmt.Sprintf("[%s:%d] [%X] [%s] %s", filename, line, log.Identifier, log.Context, format)

	fn(fmtString, v...)

}

func (log *Logger) Info(format string, v ...interface{}) *Logger {
	log.fmtLog(log.Logger.Info, format, v...)
	return log
}

func (log *Logger) Warn(format string, v ...interface{}) *Logger {
	log.fmtLog(log.Logger.Warn, format, v...)
	return log
}

func (log *Logger) Start(apiName string) *Logger {
	apiFormat := fmt.Sprint("[", apiName, "] ", "[START]")
	log.fmtLog(log.Logger.Info, apiFormat)
	return log
}

func (log *Logger) End(apiName string) {
	apiFormat := fmt.Sprint("[", apiName, "] ", "[END]")
	log.fmtLog(log.Logger.Info, apiFormat)
}