package TestCode

import "testing"

// Benchmark Example
func BenchmarkPower(b *testing.B) {
    for i:= 0; i< b.N; i++ {
        Power(4, 12)
    }
}

// go -test -run=BenchmarkPower -bench=.
/*
    goos: darwin
    goarch: amd64
    pkg: github.com/8luebottle/Play-Lab/Go/TestCode
    BenchmarkPower-12       42216885                27.3 ns/op
    PASS
    ok      github.com/8luebottle/Play-Lab/Go/TestCode      1.280s
*/