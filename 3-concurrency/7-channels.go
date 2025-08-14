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
	ch := make(chan string)
	// ["get"]
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("doing the get request and fetch the results")
		time.Sleep(time.Second * 2)
		fmt.Println("sending the result to the channel")
		ch <- "get" // send
	}()

	go func() {
		defer wg.Done()
		x := <-ch // recv // it would block until channel has not received the value
		fmt.Println("start working on the received value", x)
	}()

	wg.Wait()
}
