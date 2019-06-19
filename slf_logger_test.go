package slog

import "testing"

// BenchmarkArray1-12    	2000000000	         0.26 ns/op	       0 B/op	       0 allocs/op
func BenchmarkArray1(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var arr [1]int
		tmp := arr[:]
		tmp[0] = 100
	}
}

// BenchmarkArray1-12    	2000000000	         0.26 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSlice1(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		tmp := make([]int, 1, 1)
		tmp[0] = 100
	}
}

// BenchmarkNoLog-12    	50000000	        31.8 ns/op	      32 B/op	       1 allocs/op
func BenchmarkNoLog(b *testing.B) {
	SetLevel(LEVEL_WARN)
	log := GetLogger()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		log.Info("hahahhhhh %v", nil)
	}
}

// BenchmarkNoLog2-12    	500000000	         3.34 ns/op	       0 B/op	       0 allocs/op
func BenchmarkNoLog2(b *testing.B) {
	SetLevel(LEVEL_WARN)
	log := GetLogger()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if log.IsInfoEnabled() {
			log.Info("hahahhhhh %v", nil)
		}
	}
}

func TestLoggerFields(t *testing.T) {
	log1 := GetLogger()
	log1.BindFields(Fields{"age": 18})
	log1.Debug("hell1")

	log2 := log1.WithFields(Fields{"score": 100.0})
	log2.Info("hello2")

	log2.WithFields(Fields{"fav": "basketball"}).Warn("hello3")
}

func TestLoggerIsEnabled(t *testing.T) {
	SetLevel(LEVEL_WARN)
	l := GetLogger()
	if l.IsDebugEnabled() {
		l.Debug("debug....")
	}
	if l.IsInfoEnabled() {
		l.Info("info....")
	}
}
