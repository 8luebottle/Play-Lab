// TDD : always write the test first.
package TestCode

import "testing"

func TestIterator(t *testing.T) {
    char := "remarked"
    i := Iterator(char, 3)
    o := "remarkedremarkedremar"

    if i != o {
        t.Errorf("expected %v but got %v", i, o)
    }
}

// go test
/*
    --- FAIL: TestIterator (0.00s)
   iterator_test.go:12: expected remarkedremarkedremarked but got remarkedremarkedremar
*/