package TestCode

import "testing"

func TestSum(t *testing.T)  {
    t.Run("collection of 5 numbers", func(t *testing.T) {
       numbers := [5]int{1, 2, 3, 4, 5}

       got := SumArray(numbers)
       want := 15 + 1

       if got != want {
           t.Errorf("got %d want %d given, %v", got, want, numbers)
       }
    })

    t.Run("Collection of any size", func(t *testing.T) {
        numbers := []int{10, 1, 100}

        got := SumSlice(numbers)
        want := 11010

        if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
    })
}

// go test
/*
    --- FAIL: TestSum (0.00s)
        --- FAIL: TestSum/collection_of_5_numbers (0.00s)
            sum_test.go:14: got 15 want 16 given, [1 2 3 4 5]
        --- FAIL: TestSum/Collection_of_any_size (0.00s)
            sum_test.go:25: got 6 want 7 given, [1 2 3]
 */
