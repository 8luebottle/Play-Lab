package main

import (
	"fmt"
	"time"
)

const _MaxWorker = 4

var urls = []string{
		"https://www.github.com/8luebottle",
		"https://babytiger.netlify.app/",
		"https://www.github.com/8luebottle/TIL",
		"https://velog.io/@8luebottle/about",
	}

// jobs == sender | results == receiver
func DoWork(id int, jobs <-chan string, results chan <- string) {
	for j := range jobs {
		fmt.Println("worker", id, " started job : ", j)
		time.Sleep(time.Millisecond * 777)
		results <- j
	}
}

func main() {
	jobs := make(chan string, 100)
	r := make(chan string, 100)

	// Create Workers
	for w := 1; w <= _MaxWorker; w ++ {
		go DoWork(w, jobs, r)
	}

	// Give Jobs to Workers
	for _, job := range urls {
		jobs <- job
	}
	close(jobs)

	// Receive results
	for i := 1; i <= len(urls); i++ {
		<-r
	}
}

// *********** (Printing Results) *************//
/*
	worker 4  started job :  https://www.github.com/8luebottle
	worker 1  started job :  https://babytiger.netlify.app/
	worker 3  started job :  https://velog.io/@8luebottle/about
	worker 2  started job :  https://www.github.com/8luebottle/TIL
*/
