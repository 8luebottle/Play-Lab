package TestCode

import "testing"

// Analyzing the performance of Go functions with Benchmarks.
// Benchmark Example
func BenchmarkPower(b *testing.B) {
    for i:= 0; i< b.N; i++ {
        Power(4, 12)
    }
}

// go -test -run=Power -bench=.
/*
    goos: darwin
    goarch: amd64
    pkg: github.com/8luebottle/Play-Lab/Go/TestCode
    BenchmarkPower-12       45195823                26.5 ns/op
    PASS
    ok      github.com/8luebottle/Play-Lab/Go/TestCode      2.545s
 */
