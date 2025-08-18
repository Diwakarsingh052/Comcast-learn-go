package main

import (
	"fmt"
	"sync"
	"time"
)

// go build -race program.go // to run the program
var x int = 1

// maps are not safe for concurrent access, use mutex for safe operations
var user = map[int]string{
	1: "John",
}

func main() {
	wg := new(sync.WaitGroup)
	m := new(sync.Mutex)
	for i := 1; i <= 5; i++ {
		wg.Add(2)
		go UpdateX(i, wg, m)
		go PrintX(wg, m)
	}
	wg.Wait()
	fmt.Println("last value", x)
}

func UpdateX(val int, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	// critical section
	// this is the place where we access the shared resource

	// when a goroutine acquires a lock, another goroutine can't access the critical section
	// until the lock is not released

	// data race situations
	//	- at least one concurrent write and n number of reads
	//	- n number of concurrent writes
	// 	- n number of concurrent writes and n number of concurrent reads
	// 	Note - Data race doesn't happen if there are only concurrent reads
	m.Lock()
	defer m.Unlock()
	fmt.Println("writing", val)
	x = val
	time.Sleep(1 * time.Second)

}

func PrintX(wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	defer m.Unlock()
	fmt.Println("read", x)
	time.Sleep(4 * time.Second)
}
