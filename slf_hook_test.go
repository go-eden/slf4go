package slog

import (
	"testing"
	"time"
)

func TestRegisterHook(t *testing.T) {
	RegisterHook(func(log *Log) {
		println(log)
	})

	log := GetLogger()
	log.Trace("are you prety?", true)
	log.Debugf("are you prety? %t", true)
	log.Info("how old are you? ", nil)
	log.Infof("i'm %010d", 18)
	log.Warn("you aren't honest! ")
	log.Warnf("haha%02d %v", 1000, nil)
	log.Trace("set level to warn!!!!!")
	Trace("what?")
	log.Info("what?")
	log.Error("what?")
	log.Errorf("what?..$%s$", "XD")
	log.Fatalf("import cycle not allowed! %s", "shit...")
	log.Fatal("never reach here")
	time.Sleep(time.Millisecond * 10)
}
