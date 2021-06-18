package slog

import (
	"testing"
)

// BenchmarkNoLog-12    	50000000	        31.8 ns/op	      32 B/op	       1 allocs/op
func BenchmarkNoLog(b *testing.B) {
	SetLevel(WarnLevel)
	log := GetLogger()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		log.Info("hahahhhhh %v", nil)
	}
}

// BenchmarkNoLog2-12    	500000000	         3.34 ns/op	       0 B/op	       0 allocs/op
func BenchmarkNoLog2(b *testing.B) {
	SetLevel(WarnLevel)
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
	SetLevel(WarnLevel)
	l := GetLogger()
	if l.IsDebugEnabled() {
		l.Debug("debug....")
	}
	if l.IsInfoEnabled() {
		l.Info("info....")
	}
}

// BenchmarkLoggerIsEnabled-12    	85532004	        14.14 ns/op
func BenchmarkLoggerIsEnabled(b *testing.B) {
	SetLevel(WarnLevel)
	SetLoggerLevel("abc", InfoLevel)
	SetLoggerLevel("xyz", TraceLevel)

	l := NewLogger("abc")
	for i := 0; i < b.N; i++ {
		_ = l.IsInfoEnabled()
	}
}
