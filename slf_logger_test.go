package xlog

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
