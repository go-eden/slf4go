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

// BenchmarkLogger-12    	 3000000	       552 ns/op	     176 B/op	       2 allocs/op
func BenchmarkLogger(b *testing.B) {
	SetDriver(new(NilDriver))
	log := GetLogger()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		log.Info("hello world")
	}
}
