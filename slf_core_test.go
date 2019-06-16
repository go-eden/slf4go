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

func TestLoggerFormat(t *testing.T) {
	log := NewLogger("test")
	log.Tracef("arr: %v, %d, %s", []int{1, 2, 3}, 102, "haha")
	log.Tracef("arr: %d, %d, %f", 123, 102, 122.33)
}

/**
  BenchmarkLoggerCheckEnable-8      	500000000	         3.16 ns/op	       0 B/op	       0 allocs/op
  BenchmarkLoggerNotCheckEnable-8   	50000000	        32.9 ns/op	      16 B/op	       1 allocs/op
*/
func BenchmarkLoggerCheckEnable(b *testing.B) {
	logger := NewLogger("test")
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if logger.IsEnableTrace() {
			logger.Tracef("this is a test, b: %v, ", b)
		}
	}
}
func BenchmarkLoggerNotCheckEnable(b *testing.B) {
	logger := NewLogger("test")
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		logger.Tracef("this is a test, b: %v, ", b)
	}
}
