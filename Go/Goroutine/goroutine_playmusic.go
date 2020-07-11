// Test concurrency
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

var (
	shortSong = 10
)

var (
	wave1 = "■ 1"
	wave2 = "■■ 2"
	wave3 = "■■■ 3"
	wave4 = "■■■■ 4"
	wave5 = "■■■■■ 5"
	wave6 = "■■■■■■ 6"
)

func soundWave(s string) {
	for i := 0; i < shortSong; i++ {
		fmt.Printf("%v\n", s)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}

func playMusic() {
	wg.Add(1)
	go soundWave(wave1)
	wg.Add(1)
	go soundWave(wave2)
	wg.Add(1)
	go soundWave(wave3)
	wg.Add(1)
	go soundWave(wave4)
	wg.Add(1)
	go soundWave(wave5)
	wg.Add(1)
	go soundWave(wave6)
}

func main() {
	playMusic()
	wg.Wait()
}

// *********** (Printing Results) *************//
/*
	■■■■■■ 6
	■■■■ 4
	■■■■■ 5
	■■■ 3
	■ 1
	■■ 2
	■■ 2
	■ 1
	■■■■■■ 6
	■■■■■ 5
	■■■ 3
	■■■■ 4
	■■■■ 4
	■■■ 3
	■■ 2
	■ 1
	■■■■■■ 6
	■■■■■ 5
	■■■■■ 5
	■■ 2
	■ 1
	■■■■■■ 6
	■■■ 3
	■■■■ 4
	■■■■■■ 6
	■■■■ 4
	■■■ 3
	■ 1
	■■■■■ 5
	■■ 2
	■■ 2
	■■■ 3
	■■■■■■ 6
	■■■■ 4
	■ 1
	■■■■■ 5
	■ 1
	■■ 2
	■■■ 3
	■■■■ 4
	■■■■■■ 6
	■■■■■ 5
	■■■■■ 5
	■■■■ 4
	■■■■■■ 6
	■ 1
	■■■ 3
	■■ 2
	■ 1
	■■■■■ 5
	■■■■ 4
	■■■■■■ 6
	■■ 2
	■■■ 3
	■■■ 3
	■■■■■■ 6
	■■ 2
	■■■■■ 5
	■ 1
	■■■■ 4
*/
