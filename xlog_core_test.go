package xlog

import (
	"runtime"
	"testing"
	"time"
)

func TestGetLogger(t *testing.T) {
	logger := NewLogger()
	logger.Trace("are you prety?", true)
	logger.Debugf("are you prety? %t", true)
	logger.Info("how old are you? ", nil)
	logger.Infof("i'm %010d", 18)
	logger.Warn("you aren't honest! ")
	//logger.Warnf("haha%02d", 1000, nil)
	logger.Trace("set level to warn!!!!!")
	logger.Trace("what?")
	logger.Info("what?")
	logger.Error("what?")
	logger.Errorf("what?..$%s$", "XD")
	logger.Fatalf("import cycle not allowed! %s", "shit...")
	logger.Fatal("never reach here")
	time.Sleep(time.Millisecond * 10)
}

func TestLoggerFormat(t *testing.T) {
	logger := GetLogger("test")
	logger.Tracef("arr: %v, %d, %s", []int{1, 2, 3}, 102, "haha")
	logger.Tracef("arr: %d, %d, %f", 123, 102, 122.33)
}

/**
  BenchmarkLoggerCheckEnable-8      	500000000	         3.16 ns/op	       0 B/op	       0 allocs/op
  BenchmarkLoggerNotCheckEnable-8   	50000000	        32.9 ns/op	      16 B/op	       1 allocs/op
*/
func BenchmarkLoggerCheckEnable(b *testing.B) {
	logger := GetLogger("test")
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if logger.IsEnableTrace() {
			logger.Tracef("this is a test, b: %v, ", b)
		}
	}
}
func BenchmarkLoggerNotCheckEnable(b *testing.B) {
	logger := GetLogger("test")
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		logger.Tracef("this is a test, b: %v, ", b)
	}
}

func TestPackage(t *testing.T) {
	file, line := findCaller()
	t.Log(file, line)

	name, file, line := findCallerFunc()
	t.Log(name, file, line)
}

// BenchmarkFindCaller-12    	 3000000	       498 ns/op	     184 B/op	       2 allocs/op
func BenchmarkFindCaller(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		findCaller()
	}
}

// BenchmarkFindCallerFunc-12    	 5000000	       391 ns/op	     184 B/op	       2 allocs/op
func BenchmarkFindCallerFunc(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		findCallerFunc()
	}
}

func findCaller() (file string, line int) {
	_, file, line, _ = runtime.Caller(0)
	return
}

func findCallerFunc() (name string, file string, line int) {
	pc, file, line, _ := runtime.Caller(0)
	if f := runtime.FuncForPC(pc); f != nil {
		name = f.Name()
	}
	return
}

/**
skip=0
BenchmarkCaller-12    	 3000000	       440 ns/op	     184 B/op	       2 allocs/op
skip=1
BenchmarkCaller-12    	 3000000	       480 ns/op	     184 B/op	       2 allocs/op
skip=2
BenchmarkCaller-12    	 2000000	       636 ns/op	     184 B/op	       2 allocs/op

For better performance, should invoke caller at upper stack
*/
func BenchmarkCaller(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		runtime.Caller(2)
	}
}
