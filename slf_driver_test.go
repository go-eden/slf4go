package slog

import (
	"testing"
)

func TestNilDriver(t *testing.T) {
	d := &NilDriver{}
	SetDriver(d)

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

// use ring slice
// BenchmarkStdDriverChannel-12    	 1638012	       728.1 ns/op	     144 B/op	       3 allocs/op
// use channel again
// BenchmarkStdDriverChannel-12    	 1433360	       822.1 ns/op	     168 B/op	       4 allocs/op
func BenchmarkStdDriverChannel(b *testing.B) {
	d := newStdDriver(1 << 12)
	d.stdout = nil
	SetDriver(d)
	log := GetLogger()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Infof("hello world....%v", 1234.1)
	}
}

// --------------------------------

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
