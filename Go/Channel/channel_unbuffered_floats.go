package main

import (
	"fmt"
)

func main() {
	f := 3.48
	l := 5.99
	c := make(chan float64)  // Unbuffered float64 type channel

	go func() {
		c <- f
		c <- l
	}()

	showValue(c)
}

func showValue(channel chan float64) {
	v1 := <- channel
	v2 := <- channel
	printer(v1, v2)
}

func printer(value1, value2 float64){
	fmt.Printf("\n value1 : %v \t value2 : %v", value1, value2)
}

// ********** (Printing Results) *************//
/*
 	value1 : 3.48   value2 : 5.99
*/
