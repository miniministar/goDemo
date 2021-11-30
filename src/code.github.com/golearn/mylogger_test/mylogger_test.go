package mylogger_test

import (
	"goDemo/src/code.github.com/golearn/mylogger"
	"testing"
	"time"
)

func TestMylogger(t *testing.T) {
	log := mylogger.NewLog("debug")

	for {
		log.Trace("这是Trace日志，err: %s", "this error trace")
		log.Debug("这是Debug日志，err: %s", "this error debug")
		log.Info("这是Info日志，err: %v %d", "this error debug", 10)
		log.Warning("这是Warning日志")
		log.Error("这是Error日志")
		log.Fatal("这是Fatal日志")
		time.Sleep(time.Second * 3)
	}
}
