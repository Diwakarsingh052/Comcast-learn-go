package main

import (
	"fmt"
	"sync"
)

var x int = 1

// maps are not safe for concurrent access, use mutex for safe operations
var user = map[int]string{
	1: "John",
}

func main() {
	wg := new(sync.WaitGroup)
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go UpdateX(i, wg)
	}
	wg.Wait()
	fmt.Println("last value", x)
}

func UpdateX(val int, wg *sync.WaitGroup) {
	defer wg.Done()
	x = val
	//time.Sleep(1 * time.Second)
	fmt.Println(x)
}
