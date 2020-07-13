// Reference : https://play.golang.org/p/-wMKASb1Fh

/*
goroutine, channel

channels are to communicate with goroutines
to synchronize with goroutines(including main function, main goroutine)

We Launch the goroutine in background and Move onto the next.
We do not wait for the goroutine function to return.
(main goroutine does not block until its return)

<- ch
Receive from ‘ch’, and Discard the sent value
Wait until the channel ‘ch’ outputs some value

By default, sends and receives block until the other side is ready.
That is, every single send will block
until another goroutine receives from the channel.

This allows goroutines to synchronize without
explicit locks or condition variables.
*/
package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func main() {
	/**************************************/
	// goroutine & channel #01
	// run goroutine without time.Sleep
	// This is how goroutine communicates with "main" goroutine
	ch1 := make(chan bool)
	fmt.Println(len(ch1), cap(ch1)) // 0 0

	go func() {
		println(1)
		ch1 <- true
		fmt.Println(len(ch1), cap(ch1)) // 0 0
	}()

	// With channel, main goroutine blocks!
	<-ch1
	fmt.Println(len(ch1), cap(ch1)) // 0 0

	// The main goroutine blocks
	// until it receives a message from the shared channel.

	// We do not need time.Sleep to run goroutine
	// Wait until we retrieve from the channel ch
	// Once we receive, we discard it
	// Receivers always block until there is data to receive
	// That's why we do not proceed to the end of main
	// until we receive from channel

	// Remember by default, channel sends and receives
	// block until the other side is ready.
	// That is, every single send will block
	// until another goroutine receives from the channel.

	// That is, once we send to channel
	// and when the caller tries to read from the channel
	// , it will block until a value is sent.

	/**************************************/
	// goroutine & channel #02
	// run goroutine without time.Sleep
	m2 := make(map[int]string)
	m2[2] = "First"
	ch2 := make(chan bool)
	go func() {
		m2[2] = "Second"
		ch2 <- true
	}()

	// block until a value is sent to ch2
	// and retrieve from channel ch2
	// without this line, the output would be "First"
	_ = <-ch2          // or <-ch
	fmt.Println(m2[2]) // Second

	/**************************************/
	// goroutine & channel #03 ★★★
	// range over channel without close
	ch3 := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch3 <- i
		}
	}()

	// we don't need to close
	// since we end looping before limit
	for i := 0; i < 2; i++ {
		print(<-ch3, ",")
	}
	// 0,1,

	/*
	   Only the sender should a channel

	   Sending to a closed channel will cause a panic

	   channels are not files
	   , so we usually do not need to close them

	   Closing is only mandatory
	   when we need to tell the receiver
	   that there is no more value to come
	   like when we terminate a range loop
	*/

	/**************************************/
	// close #04
	// range over channel with close
	ch4 := make(chan int)
	go func() {
		defer close(ch4)
		for i := 10; i < 13; i++ {
			ch4 <- i
		}
	}()

	/*
	   Note that when we traverse channels,
	   the range ends when the channel is closed.
	   Make sure close the channel !!
	*/
	for v := range ch4 {
		print(v, ",")
	}
	// 10,11,12,

	println()
	/**************************************/
	// close #05 ★★★★★★★★★
	// close shut down the channel after
	// the last sent value is received
	// Synchronizing with the last value
	// can be achieved with goroutine

	// Without goroutine, we need to buffer the channel
	// Can hold 2 elements until sending blocks
	qch := make(chan int, 2)

	// without goroutine
	qch <- 10
	qch <- 20
	close(qch)
	for v := range qch {
		print(v, ",")
	}
	// 10,20,

	// qch <- 100
	// runtime error: send on closed channel

	// But receiving from closed channel is possible

	// any value received from closed channel succeeds without blocking
	// , returning the zero value for the channel element.
	println(<-qch, <-qch, <-qch, <-qch, <-qch)
	// If the channel is string type, [no output]
	// 10,20,0 0 0 0 0

	// ★★★
	x5, ok5 := <-qch
	fmt.Println(x5, ok5)
	// 0 false
	// closed channel sets ok to false

	fmt.Println(reflect.TypeOf(qch), len(qch))
	// chan int 0

	/**************************************/
	// close #06 ★★★
	cl6 := make(chan int, 2)
	cl6 <- 1
	close(cl6)
	fmt.Println("cl6:", <-cl6, <-cl6, <-cl6)
	/*
	   [Output]
	   Success 1 0 0


	   [Explain]
	   Here close is mandatory.

	   close shut down the channel after the last-sent value is received
	   Once closed, We CANNOT send on closed channels.

	   Any value received from closed channel succeeds without blocking.
	   So we need close(ch).
	   We can retrieve any values from closed channel, without blocking.
	   Once emptied, the closed channel returns the zero value.
	   The zero value of integer is 0.

	   x, ok := <-c
	   a closed channel sets ok to false
	*/

	/**************************************/
	// close #07 ★★★
	tc07 := make(chan string, 2)
	tc07 <- "A"
	close(tc07)
	fmt.Println("Success", <-tc07, <-tc07, <-tc07)
	fmt.Println(<-tc07)
	/*
	   [Output]
	   Success A

	
	   [Explain]
	   Here close is mandatory
	   because we are receiving more values than sent values

	   close shut down the channel after the last-sent value is received
	   Once closed, the channel is shut down.
	   We CANNOT send on closed channels.

	   Any value received from closed channel succeeds without blocking.
	   So we need close(ch).
	   We can retrieve any values from closed channel, without blocking.
	   Once emptied, the closed channel returns the zero value.
	   The zero value of string is null string(empty string).

	   x, ok := <-c
	   a closed channel sets ok to false
	*/

	/**************************************/
	// close #08 ★★★
	tc08 := make(chan int, 2)
	tc08 <- 1
	tc08 <- 2
	close(tc08)
	fmt.Println("Success", <-tc08, <-tc08)

	/*
	   [Output]
	   Success 1 2


	   [Explain]
	   Here close is NOT mandatory.

	   Any value received from closed channel succeeds without blocking.
	   So we need close(tc08).
	   We can retrieve any values from closed channel, without blocking.
	   Once emptied, the closed channel returns the zero value.
	   The zero value of integer is 0.

	   Here we only retrieve TWO times from the channel
	   to which only TWO values were sent.
	   So we do not get to the point that the channel is empty.
	   That is why we do not get any zero value of integer.
	*/

	/**************************************/
	// close #09 ★★★
	tc09 := make(chan int, 2)
	tc09 <- 1
	// close(tc09)
	// tc09 <- 2
	/*
		[Output]
		panic: runtime error: send on closed channel
		goroutine 1 [running]:


		[Explain]
		We CANNOT send on closed channels

		Same here
		ch := make(chan int, 2)
		ch <- 1
		ch <- 2
		close(ch)
		fmt.Println(<-ch)
		ch <- 1
	*/

	/**************************************/
	// close #10 ★★★
	tc10 := make(chan int, 2)
	tc10 <- 1
	tc10 <- 2
	close(tc10)

	for i := 0; i < cap(tc10)+1; i++ {
		v, ok := <-tc10
		fmt.Println(v, ok)
	}

	/*
	   [Output]
	   1 true
	   2 true
	   0 false


	   [Explain]
	   Here close is mandatory.
	   Without close, we have deadlock error.

	   We can use defer for close
	   or place close statement right before the function exit.
	   And the channel must be closed before being traversed with range.

	   We can retrieve any values from closed channel, without blocking.
	   Once emptied, the closed channel returns the zero value.
	   The zero value of string is null string(empty string).

	   v, ok := <-ch
	   A closed channel sets ok to false.
	*/

	/**************************************/
	// close #11 ★★★
	s11 := []int{1, 2, 3, 4, 5}
	tc11 := make(chan int)
	go func() {
		defer close(tc11)
		for _, elem := range s11 {
			tc11 <- elem
		}
	}()
	// close(tc11)
	for elem := range tc11 {
		fmt.Printf("%v ", elem)
	}

	/*
		[Output]
		1 2 3 4 5


		[Explain]
		Here close is mandatory.
		Without close, we have deadlock error.
		But close statement is not being run with goroutine.
		Then there is no synchronization with the channel and goroutines.

		Before the goroutine even started
		close(tc11) shut down the channel
		at the time that the channel is empty
		because no value has been sent yet from the goroutine.

		Before any value had ever been sent to channel from the goroutine
		the for-loop ends and the program ends.


		If we wait giving the program enough time
		to run goroutine before program exit,
		we have deadlock errors
		because during that time
		the goroutine will send values to channel
		that has already been shut down by close statement.

		And We can't send on closed channels.
	*/

	/**************************************/
	// close #12 ★★★
	tasks := make(chan int, 5)
	finished := make(chan bool)

	go func() {
		for {
			i, more := <-tasks
			if more {
				fmt.Println("Received task", i)
			} else {
				fmt.Println("Received ALL")
				finished <- true
				return
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		tasks <- i
		fmt.Println("Sent task", i)
	}
	close(tasks)
	fmt.Println("Sent ALL")

	<-finished
	close(finished)

	/*
	   Sent task 1
	   Sent task 2
	   Sent task 3
	   Sent ALL
	   Received task 1
	   Received task 2
	   Received task 3
	   Received ALL
	*/

	/**************************************/
	// close #13 ★★★
	traverse := func(arr []string) <-chan string {
		ch := make(chan string)
		go func() {
			// not here
			// close(ch)

			for index, value := range arr {
				ch <- fmt.Sprintf("INDEX %d,  VALUE %s", index, value)
			}
			// make sure to close the channel
			// when we range over channel
			// without this,
			// fatal error: all goroutines are asleep - deadlock!
			close(ch)
		}()
		return ch
	}

	lasli := []string{"A", "B", "C"}
	msgs := traverse(lasli)

	// iterate over channel, until the channel is closed.
	for value := range msgs {
		fmt.Println(value)
	}

	/*
	   INDEX 0,  VALUE A
	   INDEX 1,  VALUE B
	   INDEX 2,  VALUE C

	   we can also return channel with function
	   Note that when we traverse channels,
	   the range ends when the channel is closed.
	   Make sure close the channel !!
	*/

	/**************************************/
	// close #14 ★★★
	// Create a new channel with `make(chan val-type)`.
	// Channels are typed by the values they convey.
	pmsg := make(chan string, 5)

	// send a value to channel message
	go func() {
		// make sure to close the channel
		// when we range over channel
		// without this
		// Error : goroutine 1 [chan receive]:
		defer close(pmsg)
		pmsg <- "ping"
		pmsg <- "pong1"
		pmsg <- "pong2"
		pmsg <- "pong3"
	}()

	for v := range pmsg {
		fmt.Println(v)
	}
	/*
	   ping
	   pong1
	   pong2
	   pong3

	   Without close, we have deadlock error.

	   We can use defer for close
	   or place close statement right before the function exit.
	   And the channel must be closed before being traversed with range.

	   defer pushed a function call onto a list
	   After the surrounding function returns,
	   the list is to be executed in LIFO(Last In First Out) order

	   The surrounding function of defer close(ch) is the goroutine

	   Note that when we traverse channels,
	   the range ends when the channel is closed.
	   Make sure to close the channel before traversing !!
	*/

	/**************************************/
	// close #15 ★★★
	dc := make(chan bool)

	aslice := []string{"a", "b", "c"}
	for _, v := range aslice {
		go func() {
			fmt.Println(v)
			dc <- true
		}()
	}

	// wait for all goroutines to complete before exiting
	for _ = range aslice {
		<-dc
	}
	/*
	   c
	   c
	   c
	*/

	/**************************************/
	// channel type #01 ★★★
	// send-only
	rc1 := make(chan<- int)
	go func() {
		rc1 <- 1
	}()
	fmt.Println(reflect.TypeOf(rc1), len(rc1), cap(rc1))
	// chan<- int 0 0

	// <-rc1
	// invalid operation: <-rc1 (receive from send-only type chan<- int)

	// rc1 <- 2
	// fatal error: all goroutines are asleep - deadlock!
	// since it is not buffered
	// we need to use goroutine

	// close(rc1)
	// x6, ok6 := <-rc1
	// <-rc1 (receive from send-only type chan<- int)

	/**************************************/
	// channel type #02 ★★★
	// receive-only
	// Frequently used as function return type
	rc2 := make(<-chan int)
	go func() {
		// rc2 <- 1
		// invalid operation: rc2 <- 1 (send to receive-only type <-chan int)
	}()
	fmt.Println(reflect.TypeOf(rc2), len(rc2), cap(rc2))
	// <-chan int 0 0

	// <-rc2
	// fatal error: all goroutines are asleep - deadlock!

	// rc2 <- 2
	// invalid operation: rc2 <- 2 (send to receive-only type <-chan int)

	// close(rc2)
	// invalid operation: close(rc2) (cannot close receive-only channel)
	// close executed only by the sender, never by the receiver
	// Shut down the channel after the last sent value is received
	// We CAN'T send to receive-only channel
	// So close is not possible

	sendNumbers := func() <-chan int {
		// make bi-directional channel
		ch := make(chan int)

		go func() {
			defer close(ch)
			for i := 0; i < 3; i++ {
				ch <- i
				fmt.Println("ch:", reflect.TypeOf(ch), len(ch), cap(ch))
				// ch: chan int 0 0
			}
		}()
		fmt.Println("ch:", reflect.TypeOf(ch), len(ch), cap(ch))
		// ch: chan int 0 0

		// will be converted to receive-only channel
		return ch
	}

	ch7 := sendNumbers()
	fmt.Println("ch7:", reflect.TypeOf(ch7), len(ch7), cap(ch7))
	// ch7: <-chan int 0 0

	// now we retrieve values from receive-only channel ch7
	for v := range ch7 {
		fmt.Printf("%d ", v)
	}
	// 0 1 2

	/**************************************/
	// channel type #03 ★★★
	// interface
	mt := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
	}
	lslice := []string{"A", "B", "C"}

	type Person struct {
		name string
		age  int
	}

	// main.Person
	p := Person{
		name: "Lee",
		age:  10,
	}

	// *main.Person
	pt := new(Person)
	pt.name = "Google"
	pt.age = 20

	tsl := []interface{}{1, "A", mt, lslice, p, pt}
	for key, value := range tsl {
		fmt.Printf("%+v: %+v (%+v) \n", key, value, reflect.TypeOf(value))
	}
	/*
	   0: 1 (int)
	   1: A (string)
	   2: map[H:Hydrogen He:Helium] (map[string]string)
	   3: [A B C] ([]string)
	   4: {name:Lee age:10} (main.Person)
	   5: &{name:Google age:20} (*main.Person)
	*/

	chtt := make(chan interface{})
	go func() {
		defer close(chtt)
		// traverse slice
		// and send each element to the channel chtt
		for _, value := range tsl {
			chtt <- value
		}
	}()

	// traverse slice
	// and retrieve each value from the channel chtt
	for value := range chtt {
		fmt.Println("channel:", value)
	}
	/*
	   channel: 1
	   channel: A
	   channel: map[H:Hydrogen He:Helium]
	   channel: [A B C]
	   channel: {Lee 10}
	   channel: &{Google 20}
	*/

	/**************************************/
	// channel type #04 ★★★
	pkFunc := func(fns ...func()) func() {
		return fns[rand.Intn(len(fns))]
	}

	pdFunc := func(c chan func(), n int, fns ...func()) {
		defer close(c)
		for i := 0; i < n; i++ {
			c <- pkFunc(fns...)
		}
	}

	rand.Seed(time.Now().Unix())

	x := 10
	fns := []func(){
		func() { x += 1 },
		func() { x -= 1 },
		func() { x *= 2 },
		func() { x /= 2 },
		func() { x *= x },
	}

	cf := make(chan func())
	go pdFunc(cf, 10, fns...)

	for fn := range cf {
		fn()
		print(x, ",")
		time.Sleep(time.Millisecond)
	}
	// 9,18,19,9,18,324,648,647,418609,837218,

	/**************************************/
	// channel type #05 ★★★
	// empty struct, more idiomatic than boolean channel
	/*
	   http://blog.carlsensei.com/post/72359081647
	   <-chan struct{}

	   what is <-chan struct{}? type struct {} is a type for hanging methods
	   off of with no associated data. chan struct{} is a channel that passes
	   no data. You could also write chan bool but that has two potential values
	   , true or false. struct{} lets your readers know, “There’s no data
	   being passed here.” The only reason for the channel to exist is to
	   synchronize things. No other information is being sent. The arrow in
	   front of <-chan struct{} further specifies that the caller is the one
	   who will be in charge of the quit channel, not the function.
	*/
	// we can also use quit <-chan struct{}
	oDD := func(from, to int, quit chan struct{}) <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			if from%2 == 0 {
				from += 1
			}
			for i := from; i < to; i += 2 {
				select {
				case <-quit:
					return
					// exit the goroutine
				case c <- i:
				}
			}
		}()
		return c
	}
	// define as bi-directional
	quit := make(chan struct{})
	defer close(quit)
	// we can't send onto closed channel
	// but we can always retrieve from closed channel
	for v := range oDD(0, 10, quit) {
		fmt.Print(v, ", ")
		if v > 5 {
			fmt.Println("Done!")
			// return
			// or
			quit <- struct{}{}
			// necessary to signal
			// we can stop all of the go-routines that are
			// listening to the same quit channel by
			// just doing close(quit) once
		}
	}
	// 1, 3, 5, 7, Done!
}
