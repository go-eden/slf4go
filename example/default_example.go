package main

import (
	log "github.com/go-eden/slf4go"
	"time"
)

func main() {
	log.Trace("trace log")
	log.Tracef("trace time: %v", time.Now())
	log.Debug("debug log")
	log.Debugf("debug time: %v", time.Now())
	log.Info("info log")
	log.Infof("info log: %v", time.Now())
	log.Warn("warn log")
	log.Warnf("warn log: %v", time.Now())
	log.Error("error log")
	log.Errorf("error time: %v", time.Now())
	log.Panic("panic log")
	log.Panicf("panic time: %v", time.Now())
	log.Fatal("fatal log")
	log.Fatalf("fatal time: %v", time.Now())
}
