package TestCode

import (
	"fmt"
	"strconv"
)

// EntryPoint
func main() {
	fmt.Println(Adder())
}

// Adder adds two integer and return as a string.
func Adder() string {
	r := 5 + 20
	stringNumb := strconv.Itoa(r)
	return fmt.Sprintf("adder : %v", stringNumb)
}
