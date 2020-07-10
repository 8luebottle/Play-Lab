package main

import (
	"fmt"
)

func main() {
	a := 2019
	b := 2020

	cIn := make(chan int, 2)
	cIn <- a
	cIn <- b

	addValue(cIn)
}

func addValue(in chan int) {
	var sum int
	for i := 0; i <= len(in); i++ {
		sum += <- in
		printer(i, sum)
	}
}

func printer(i int, sum int) {
	fmt.Printf("\n %d sum : %v\n", i, sum)
}

// ********** (Printing Results) *************//
/*
 	0 sum : 2019

 	1 sum : 4039
*/

