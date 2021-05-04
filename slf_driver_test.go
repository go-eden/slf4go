package slog

import (
	"testing"
)

type NilDriver struct {
}

func (*NilDriver) Name() string {
	return ""
}

func (*NilDriver) Print(_ *Log) {
}

func (*NilDriver) GetLevel(_ string) Level {
	return TraceLevel
}

func TestNilDriver(t *testing.T) {
	SetDriver(new(NilDriver))

	log := GetLogger()
	log.Info("what???")
}

// BenchmarkLogger-12    	 3000000	       420 ns/op	     112 B/op	       2 allocs/op
func BenchmarkLogger(b *testing.B) {
	SetDriver(new(NilDriver))

	log := GetLogger()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		log.Infof("hello world")
	}
}

// BenchmarkDefaultLogger-12    	 2475452	       419.4 ns/op	     128 B/op	       2 allocs/op
func BenchmarkDefaultLogger(b *testing.B) {
	SetDriver(new(NilDriver))

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Infof("hello world")
	}
}

// BenchmarkLoggerIsEnable-12    	201168616	         5.842 ns/op	       0 B/op	       0 allocs/op
func BenchmarkLoggerIsEnable(b *testing.B) {
	SetDriver(new(NilDriver))
	log := GetLogger()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		log.IsInfoEnabled()
	}
}
