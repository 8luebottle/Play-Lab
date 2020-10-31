package TestCode

import (
	"fmt"
	"strconv"
)

func main() {
    fmt.Printf(Adder())
}

func Adder() string {
	sum := 5 + 20
	strNumb := strconv.Itoa(sum)
	return fmt.Sprintf("Sum == %v", strNumb)
}
