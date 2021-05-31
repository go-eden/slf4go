package main

import (
	slog "github.com/go-eden/slf4go"
	"time"
)

func main() {
	//log.Trace("trace log")
	//log.Tracef("trace time: %v", time.Now())
	//log.Debug("debug log")
	slog.Debugf("debug time: %v", time.Now())
	//log.Info("info log")
	//log.Infof("info log: %v", time.Now())
	slog.Warn("warn log")
	//log.Warnf("warn log: %v", time.Now())
	slog.Error("error log")
	//log.Errorf("error time: %v", time.Now())
	//log.Panic("panic log")
	slog.Panicf("panic time: %v", time.Now())
	//log.Fatal("fatal log")
	//log.Fatalf("fatal time: %v", time.Now())

	time.Sleep(time.Millisecond)
}
