package main

import "testing"

var result interface{}

func Benchmark_IOTA(b *testing.B) {
	var r interface{}

	for i := 0; i < b.N; i++ {
		r = Run()
	}
	result = r
}
