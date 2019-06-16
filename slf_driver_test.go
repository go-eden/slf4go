package xlog

import "testing"

type NilDriver struct {
}

func (*NilDriver) Name() string {
	return ""
}

func (*NilDriver) Print(l *Log) {
}

func (*NilDriver) GetLevel(logger string) Level {
	return LEVEL_TRACE
}

func TestNilDriver(t *testing.T) {
	SetDriver(new(NilDriver))

	log := GetLogger()
	log.Info("what???")
}

// BenchmarkLogger-12    	 3000000	       446 ns/op	     192 B/op	       2 allocs/op
func BenchmarkLogger(b *testing.B) {
	SetDriver(new(NilDriver))

	log := GetLogger()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		log.Info("hello world")
	}
}

// BenchmarkDefaultLogger-12    	 3000000	       448 ns/op	     192 B/op	       2 allocs/op
func BenchmarkDefaultLogger(b *testing.B) {
	SetDriver(new(NilDriver))

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Info("hello world")
	}
}

// BenchmarkLoggerIsEnable-12    	2000000000	         1.56 ns/op	       0 B/op	       0 allocs/op
func BenchmarkLoggerIsEnable(b *testing.B) {
	SetDriver(new(NilDriver))
	log := GetLogger()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		log.IsInfoEnabled()
	}
}
