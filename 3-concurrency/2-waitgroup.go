package main

import (
	"fmt"
	"sync"
)

func main() {

	//var wg *sync.WaitGroup  // nil
	//wg := &sync.WaitGroup{} // {0,0,0} x80
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go workV2(wg)
	fmt.Println("some more main related work going on")
	wg.Wait() // it would block the main, until the counter is not back to 0
	fmt.Println("main function finished")
}

func workV2(wg *sync.WaitGroup) {
	defer wg.Done() // even your program panics, defer would run
	// defer runs when the function returns
	fmt.Println("doing some work")
}
