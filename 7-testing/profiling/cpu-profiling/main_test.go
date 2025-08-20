package main

import "testing"

var moby = `
The Project Gutenberg EBook of Moby Dick; or The Whale, by Herman
Melville

This eBook is for the use of anyone anywhere at no cost and with almost
no restrictions whatsoever.  You may copy it, give it away or re-use it
under the terms of the Project Gutenberg License included with this
eBook or online at www.gutenberg.org

`

// //go test -run none -bench . -benchtime 3s -benchmem -v -cpuprofile p.out
// go tool pprof p.out
// list
// weblist all commands would be the same as memory profiling
func BenchmarkUniqueSubstrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UniqueSubstrings(moby)
	}
}

func BenchmarkOptimizedUniqueSubstrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OptimizedUniqueSubstrings(moby)
	}
}
