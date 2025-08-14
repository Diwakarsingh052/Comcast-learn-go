package main

import (
	"fmt"
	"sync"
	"time"
)

// A send on an unbuffered channel can proceed if a receiver is ready.
// A send on a buffered channel can proceed if there is room in the buffer.

func main() {
	wg := new(sync.WaitGroup)
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("sending the values")
		ch <- 10
		ch <- 20
		ch <- 30 // in case of unbuffered this program would deadlock,
		// because there is no recv for the third value
		fmt.Println("sent the values")
	}()
	time.Sleep(time.Second * 2)
	fmt.Println("receiving the values")
	//fmt.Println(<-ch) // receiver always blocks
	//fmt.Println(<-ch)
	for i := 0; i < 2; i++ {
		fmt.Println(<-ch)
	}
	wg.Wait()
}
