package log

import (
	"runtime"
	"testing"
)

func TestPCStack(t *testing.T) {
	pc := make([]uintptr, 1, 1)
	_ = runtime.Callers(1, pc)

	t.Log(ParseStack(pc[0]))
	t.Log(ParseStack(pc[0]))
	t.Log(ParseStack(pc[0]))
}

// BenchmarkParseStack-12    	300000000	         5.83 ns/op	       0 B/op	       0 allocs/op
func BenchmarkParseStack(b *testing.B) {
	pc := make([]uintptr, 1, 1)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = runtime.Callers(1, pc)
		ParseStack(pc[0])
	}
}

// BenchmarkParseStack2-12    	 5000000	       245 ns/op	     248 B/op	       3 allocs/op
func BenchmarkParseStack2(b *testing.B) {
	pc := make([]uintptr, 1, 1)
	_ = runtime.Callers(1, pc)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parseStack(pc[0])
	}
}

// skip=1
// BenchmarkCaller-12    	10000000	       180 ns/op	       0 B/op	       0 allocs/op
// skip=2
// BenchmarkCaller-12    	10000000	       236 ns/op	       0 B/op	       0 allocs/op
// skip=3
// BenchmarkCaller-12    	 5000000	       325 ns/op	       0 B/op	       0 allocs/op
// For better performance, should invoke caller at upper stack
func BenchmarkCaller(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var pc [1]uintptr
		runtime.Callers(3, pc[:])
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
	_, file, line, _ = runtime.Caller(2)
	return
}

func findCallerFunc() (name string, file string, line int) {
	pc, file, line, _ := runtime.Caller(2)
	if f := runtime.FuncForPC(pc); f != nil {
		name = f.Name()
	}
	return
}

// BenchmarkCallers-12    	10000000	       177 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCallers-12    	 5000000	       228 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCallers-12    	 3000000	       435 ns/op	     176 B/op	       1 allocs/op
func BenchmarkCallers(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	rpc := make([]uintptr, 1, 1)
	for i := 0; i < b.N; i++ {
		runtime.Callers(2, rpc)
		_, _ = runtime.CallersFrames(rpc).Next()
	}
}
