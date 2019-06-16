package xlog

import (
	"testing"
	"time"
)

func TestGetLogger(t *testing.T) {
	log := GetLogger()
	log.Trace("are you prety?", true)
	log.Debugf("are you prety? %t", true)
	log.Info("how old are you? ", nil)
	log.Infof("i'm %010d", 18)
	log.Warn("you aren't honest! ")
	log.Warnf("haha%02d %v", 1000, nil)
	log.Trace("set level to warn!!!!!")
	log.Trace("what?")
	log.Info("what?")
	log.Error("what?")
	log.Errorf("what?..$%s$", "XD")
	log.Fatalf("import cycle not allowed! %s", "shit...")
	log.Fatal("never reach here")
	time.Sleep(time.Millisecond * 10)
}

func TestDefaultLogger(t *testing.T) {
	Trace("are you prety?", true)
	Debugf("are you prety? %t", true)
	Info("how old are you? ", nil)
	Infof("i'm %010d", 18)
	Warn("you aren't honest! ")
	Warnf("haha%02d %v", 1000, nil)
	Trace("set level to warn!!!!!")
	Trace("what?")
	Info("what?")
	Error("what?")
	Errorf("what?..$%s$", "XD")
	Fatalf("import cycle not allowed! %s", "shit...")
	Fatal("never reach here")
}
