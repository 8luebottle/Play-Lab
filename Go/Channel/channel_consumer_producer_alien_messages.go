package main

import (
	"fmt"
)

// This is a producer-consumer pattern
func main() {
	alienMessages := []string{
		"BIG BROTHER IS WATCHING YOU",
		"Perhaps one did not want to be loved so much as to be understood.",
		"War is peace. Freedom is slavery. Ignorance is strength.",
		"If you want to keep a secret, you must also hide it from yourself.",
		"But if thought corrupts language, language can also corrupt thought.",
	}

	telepathy := make(chan string)
	done := make(chan bool)

	go producer(telepathy, alienMessages)
	go consumer(telepathy, done)
	<-done
}

func producer(telepathy chan<- string, alienMessages []string) {
	for _, m := range alienMessages {
		telepathy <- m
	}
	close(telepathy) // If you miss this part go will return fatal error : all goroutines are asleep - deadlock!
}

func consumer(telepathy <-chan string, done chan<- bool) {
	for d := range telepathy {
		fmt.Printf("\nMessage From Alien : '%v'\n", d)
	}
	done <- true // If you miss this part go will return fatal error : all goroutines are asleep - deadlock!
}

// ********** (Printing Results) *************//
/*
	Message From Alien : 'BIG BROTHER IS WATCHING YOU'

	Message From Alien : 'Perhaps one did not want to be loved so much as to be understood.'

	Message From Alien : 'War is peace. Freedom is slavery. Ignorance is strength.'

	Message From Alien : 'If you want to keep a secret, you must also hide it from yourself.'

	Message From Alien : 'But if thought corrupts language, language can also corrupt thought.'
*/
