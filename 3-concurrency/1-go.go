package main

import (
	"fmt"
	"time"
)

func main() {
	go work() // the main function is spinning up this goroutine
	fmt.Println("main function finished")
	time.Sleep(time.Second * 1) // this is not recommended
}

func work() {
	fmt.Println("doing some work")
}
