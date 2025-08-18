package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := new(sync.WaitGroup)
	get := make(chan string)
	post := make(chan string)
	put := make(chan string)

	wg.Add(3)
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		get <- "get done"
	}()
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		post <- "post done"
	}()
	go func() {
		defer wg.Done()
		put <- "put done"
	}()

	//fmt.Println("get:", <-get)
	//fmt.Println("post:", <-post)
	//fmt.Println("put:", <-put)

	for i := 1; i <= 3; i++ {

		select {
		//whichever case is not blocking exec that first
		//whichever case is ready first, exec that.
		// possible cases are chan recv , send , default
		case x := <-get:
			fmt.Println("received", x)
		case x := <-post:
			fmt.Println("received", x)
		case x := <-put:
			fmt.Println("received", x)
		}
	}

	wg.Wait()
}
