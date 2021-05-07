package slog

import (
	"testing"
)

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

// when use channel(lots of [WARNING: log lost, channel is full])
// BenchmarkStdDriverChannel-12    	 1389292	       923.8 ns/op	     277 B/op	       9 allocs/op
// BenchmarkStdDriverChannel-12    	 1507141	       793.6 ns/op	     256 B/op	       8 allocs/op
// when bufsize==1M(no log lost)
// BenchmarkStdDriverChannel-12    	 1789503	       655.5 ns/op	     236 B/op	       7 allocs/op
// when asyncPrint do nothing
// BenchmarkStdDriverChannel-12    	 1638012	       728.1 ns/op	     144 B/op	       3 allocs/op
func BenchmarkStdDriverChannel(b *testing.B) {
	d := newStdDriver(1 << 12).(*StdDriver)
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
