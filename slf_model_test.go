package xlog

import (
	"github.com/huandu/go-tls"
	"os"
	"runtime"
	"testing"
	"time"
)

func TestGid(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			id := tls.ID() // Get a unique ID for current goroutine. It's guaranteed to be unique.
			t.Log(id)
		}()
	}
	time.Sleep(time.Second)
}

// BenchmarkGid-12    	100000000	        16.7 ns/op	       0 B/op	       0 allocs/op
func BenchmarkGid(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tls.ID()
	}
}

func TestPid(t *testing.T) {
	t.Log(os.Getpid())
	time.Sleep(time.Minute)
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
	l := NewLog(LEVEL_TRACE, pc[0], "")
	t.Log(l)
}

// BenchmarkNewLog-12    	10000000	       169 ns/op	     160 B/op	       1 allocs/op
// BenchmarkNewLog-12    	 2000000	       769 ns/op	     408 B/op	       4 allocs/op
// BenchmarkNewLog-12    	 2000000	       720 ns/op	     392 B/op	       4 allocs/op
// after optimization by ParseStack
// BenchmarkNewLog-12    	 3000000	       408 ns/op	     160 B/op	       1 allocs/op
// ignore Callers
// BenchmarkNewLog-12    	10000000	       145 ns/op	     160 B/op	       1 allocs/op
func BenchmarkNewLog(b *testing.B) {
	var pc [1]uintptr
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = runtime.Callers(2, pc[:])
		NewLog(LEVEL_TRACE, pc[0], "")
	}
}

func TestLevel(t *testing.T) {
	t.Log(LEVEL_FATAL.String())
	t.Log(LEVEL_TRACE)
}
