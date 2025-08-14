package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	// it is a func without a name
	// and it must be called immediately after it body ends

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("calling workV3 concurrently")
		workV3()
	}() // () this how we call an anonymous func

	wg.Wait()
	fmt.Println("calling workV3 normally")
	workV3()
}

func workV3() {
	fmt.Println("doing some work, concurrently or normally")
}
