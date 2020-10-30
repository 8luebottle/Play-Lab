package TestCode

import "testing"

func TestAdder(t *testing.T) {
	in := Adder()
	out := "adder : 26"

	if in != out {
		t.Errorf(
			"input value does not equal to output value\nInput : %v\tOutput: %v",
			in, out,
		)
	}
}
