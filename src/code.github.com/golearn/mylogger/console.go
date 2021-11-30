package mylogger

import (
	"fmt"
	"google.golang.org/protobuf/internal/errors"
	"strings"
	"time"
)

type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

//往终端里写日志相关内容
type Logger struct {
	Level LogLevel
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

func (l Logger) Trace(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [TRACE] %s", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Debug(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [DEBUG] %s", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Info(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [INFO] %s", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Warning(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [WARNING] %s", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Error(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [ERROR] %s", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Fatal(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [FATAL] %s", now.Format("2006-01-02 15:04:05"), msg)
}
