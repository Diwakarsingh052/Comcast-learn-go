package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := new(sync.WaitGroup)
	wgTask := new(sync.WaitGroup)
	get := make(chan string)
	post := make(chan string)
	put := make(chan string)
	done := make(chan struct{})

	wgTask.Add(3)
	go func() {
		defer wgTask.Done()
		time.Sleep(3 * time.Second)
		get <- "get done"
	}()
	go func() {
		defer wgTask.Done()
		time.Sleep(1 * time.Second)
		post <- "post done"
	}()
	go func() {
		defer wgTask.Done()
		put <- "put done"
		put <- "put 2 done"
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		wgTask.Wait()
		// we will close the done channel when all goroutines are finished
		//close is also a send signal, which can be received in select
		close(done)

	}()

	//fmt.Println("get:", <-get)
	//fmt.Println("post:", <-post)
	//fmt.Println("put:", <-put)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {

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
			case <-done:
				fmt.Println(" all goroutines finished")
				return
			default:
				fmt.Println("waiting...")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	wg.Wait()
}
