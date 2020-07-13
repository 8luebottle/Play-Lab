// Reference : Mr.Waggel's Blog
// Sourcecode : https://mrwaggel.be/post/golang-creating-a-worker-thread-pool-using-maps/
package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Make the workthread globally accessible
var Workpool WorkThread

type WorkThread struct {
	Todo    map[int]string
	LastKey int
	KeyLock sync.Mutex
	MapLock sync.Mutex
}

func main() {
	//Initialize the work thread
	Workpool = *new(WorkThread)
	Workpool.Todo = make(map[int]string)
	Workpool.LastKey = 0
	go Workpool.Start() //Start the work pool
	// Wait and add items to the workpool
	// Wait and add items to the workpool, simulating
	time.Sleep(3 * time.Second)
	Workpool.AddToQueue("Hash me")
	Workpool.AddToQueue("And me too")
	Workpool.AddToQueue("Don't forget to hash this one")
	time.Sleep(5 * time.Second)
	Workpool.AddToQueue("More work!")
	Workpool.AddToQueue("Not over yet")
	time.Sleep(4 * time.Second)
	Workpool.AddToQueue("Almost there")
	Workpool.AddToQueue("Last one")
	// Make sure main doesn't return before the work is done
	for {
		time.Sleep(time.Second)
	}
}

func (w *WorkThread) Start() {
	fmt.Println("Start worker")
	// This thread will wait for work
	for {
		// Check if there is a job to do
		if len(w.Todo) == 0 {
			// Nothing to do, wait for work
			fmt.Println("No jobs, waiting...")
			time.Sleep(2 * time.Second)
			continue // Move back to the beginning of the loop
		}
		// We have work to do, lets get the key of the job
		JobKey := w.GetNextKey()
		// Generate a hash with the key that holds the string,
		HashedString, _ := bcrypt.GenerateFromPassword([]byte(w.Todo[JobKey]), 15)
		fmt.Println("Job completed: " + string(HashedString))
		// Remove this string from the job queue
		w.RemoveFromQue(JobKey)
	}
}

func (w *WorkThread) Newkey() int {
	// This is to get a new key for the map
	w.KeyLock.Lock()         // Lock the LastKey integer
	defer w.KeyLock.Unlock() // Tell go to call this function when NewKey() returns
	w.LastKey++              // Increment the last used key by one to add to our map
	return w.LastKey
}

func (w *WorkThread) AddToQueue(Newstring string) {
	// Add job to the queue
	NewKey := w.Newkey()       // First get a new key
	w.KeyLock.Lock()           // Lock the map, make it safe to write too
	defer w.KeyLock.Unlock()   // Unlock the map after function return
	w.Todo[NewKey] = Newstring // Add the string with the new key to the map
	fmt.Println("Job added!")
	return
}

func (w *WorkThread) RemoveFromQue(DeletKey int) {
	w.MapLock.Lock()         // Lock the map
	defer w.MapLock.Unlock() // Call this function on return
	delete(w.Todo, DeletKey) // Remove the job from queue
	return
}

func (w *WorkThread) GetNextKey() int {
	var ReturnKey int
	// Get the first key in the map
	for key, _ := range w.Todo {
		ReturnKey = key
		break
	}
	return ReturnKey
}

// ******************* (Printing Results) ********************//
/*
	Start worker
	No jobs, waiting...
	No jobs, waiting...
	Job added!
	Job added!
	Job added!
	Job completed: $2a$15$/.ZuT3WSkEdDwJ6TSVeIouY3i7VjJofGG/2m9K6nkaddGUcfom5H.
	Job completed: $2a$15$6MtoPD6AVxfMTTTXPq49G.7u4jPzH7CTnoLExeHRZcC0KLTN66as.
	Job added!
	Job added!
	Job completed: $2a$15$AnIOSOnY5MmalOUGGXo1fuy1StuuUiMFwoyItfpcsst0qIK9iBXzG
	Job completed: $2a$15$5av1RFJ0KuY9ojXPsLLxNeuXWzu0o4.elS6pZd0H6mn9tD.Mnk4Q2
	Job added!
	Job added!
	Job completed: $2a$15$QMTCjjZg5FGr5lZVdU37pevjx5UnjV2ZRBnTtlMVcppweg.sszn0a
	Job completed: $2a$15$N3pO.3EIzW/qlXu3tEcd9ONmJUlItwEJUDGV/PihkTCt0qRngXOVq
	Job completed: $2a$15$ROvkH1COzKEUGVKl3br/fe1N9LtMn17qrxwvpPfFgyl7gCqo7XQJK
	No jobs, waiting...
	No jobs, waiting...
*/
