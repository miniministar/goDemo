package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	TRACE
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

//往终端里写日志相关内容
type Logger struct {
	Level   LogLevel
	logChan chan *LogMsg
}

type LogMsg struct {
	Level     LogLevel
	Msg       string
	FuncName  string
	FileName  string
	Timestamp string
	line      int
}

//Logger 构造夯实
func NewLog(levelStr string) Logger {
	level, e := parseLogLevel(levelStr)
	if e != nil {

	}
	return Logger{
		Level: level,
	}
}

func parseLogLevel(levelStr string) (LogLevel, error) {
	s := strings.ToLower(levelStr)
	switch s {
	case "trace":
		return TRACE, nil
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	split := strings.Split(funcName, ".")
	funcName = split[len(split)-1]
	return
}

func (l Logger) Trace(format string, a ...interface{}) {
	if l.enabled(TRACE) {
		log(TRACE, format, a...)
	}
}

func log(level LogLevel, format string, a ...interface{}) {
	var msg string
	if a != nil && len(a) > 0 {
		msg = fmt.Sprintf(format, a...)
	} else {
		msg = format
	}
	funcName, fileName, lineNo := getInfo(3)
	now := time.Now()
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLevelString(level), fileName, funcName, lineNo, msg)
}

func getLevelString(level LogLevel) string {
	switch level {
	case TRACE:
		return "TRACE"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

func (l Logger) Debug(format string, a ...interface{}) {
	if l.enabled(DEBUG) {
		log(DEBUG, format, a...)
	}
}

func (l Logger) Info(format string, a ...interface{}) {
	if l.enabled(INFO) {
		log(INFO, format, a...)
	}
}

func (l Logger) Warning(format string, a ...interface{}) {
	if l.enabled(WARNING) {
		log(WARNING, format, a...)
	}
}

func (l Logger) Error(format string, a ...interface{}) {
	if l.enabled(ERROR) {
		log(ERROR, format, a...)
	}
}

func (l Logger) Fatal(format string, a ...interface{}) {
	if l.enabled(FATAL) {
		log(FATAL, format, a...)
	}
}

func (l Logger) enabled(level LogLevel) bool {
	return level >= l.Level
}
