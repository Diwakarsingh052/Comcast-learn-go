package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	ch := make(chan int, 10)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch) // sends a signal to stop the range
		// close signal range that no more values be sent
		//and it can stop after receiving remaining values
		// once the channel is closed, we can't send more values to it
	}()

	// it would run infinitely, channel needs to be closed to stop this range
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			time.Sleep(time.Millisecond * 50)
			fmt.Println(v, "1st range")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			
			time.Sleep(time.Millisecond * 50)
			fmt.Println(v, "2nd range")
		}
	}()

	wg.Wait()

}
