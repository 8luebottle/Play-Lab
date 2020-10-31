package TestCode

import "testing"

func TestAdder(t *testing.T) {
    i := Adder()
    o := "Sum == 26"

    if i != o {
        t.Errorf("output(%v) should be same as (%v)", o, i)
    }
}

// go test
/*
    --- FAIL: TestAdder (0.00s)
   adder_test.go:10: output(Sum == 26) should be same as (Sum == 25)
*/