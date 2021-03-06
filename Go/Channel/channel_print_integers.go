package main

import (
	"fmt"
)

func main() {
	a := 4
	b := 5

	cIn := make(chan int, 2)  // Channel by default is pointer
	
	// Write data (channel <-) to the Channel cIn == push to channel
	cIn <- a               
	cIn <- b                  

	getChan(cIn)
}

func getChan(in chan int) {
	for i := 0; i <= len(in); i++ {
		// Read data (<- channel) == pop from channel
		out := <- in
		printer(in, out)
	}
}

func printer(in chan int, out int) {
	fmt.Printf("\n in : %v  out : %v", in, out)
	fmt.Printf("\tin's type: %T | out's type : %T\n", in, out)
}
// ********** (Printing Results) *************//
/*
 in : 0xc0000b8000  out : 4     in's type: chan int | out's type : int

 in : 0xc0000b8000  out : 5     in's type: chan int | out's type : int
*/

