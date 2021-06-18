package slog

import (
	"os"
	"runtime"
	"testing"
	"time"
)

func TestPid(t *testing.T) {
	t.Log(os.Getpid())
	time.Sleep(time.Second)
}

// BenchmarkPid-12    	100000000	        17.9 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPid(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		os.Getpid()
	}
}

func TestNewLog(t *testing.T) {
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l := NewLog(TraceLevel, pc[0], nil, nil, nil, nil, nil)
	t.Log(l)
}

// BenchmarkNewLog-12    	10000000	       169 ns/op	     160 B/op	       1 allocs/op
// BenchmarkNewLog-12    	 2000000	       769 ns/op	     408 B/op	       4 allocs/op
// BenchmarkNewLog-12    	 2000000	       720 ns/op	     392 B/op	       4 allocs/op
// after optimization by ParseStack
// BenchmarkNewLog-12    	 5000000	       387 ns/op	     176 B/op	       1 allocs/op
// BenchmarkNewLog-12    	 5000000	       376 ns/op	     160 B/op	       1 allocs/op
// BenchmarkNewLog-12    	 5000000	       363 ns/op	      96 B/op	       1 allocs/op
func BenchmarkNewLog(b *testing.B) {
	var pc [1]uintptr
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = runtime.Callers(2, pc[:])
		NewLog(TraceLevel, pc[0], nil, nil, nil, nil, nil)
	}
}
