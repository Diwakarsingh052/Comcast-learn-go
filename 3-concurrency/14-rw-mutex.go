package main

import (
	"fmt"
	"sync"
	"time"
)

// read write mutex

var y int = 1

func main() {
	wg := new(sync.WaitGroup)
	m := new(sync.RWMutex)
	for i := 1; i <= 5; i++ {
		wg.Add(2)
		go UpdateY(i, wg, m)
		go PrintY(wg, m)
	}
	wg.Wait()
	fmt.Println("last value", y)
}

func UpdateY(val int, wg *sync.WaitGroup, m *sync.RWMutex) {
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
	m.Lock() // when Write lock is acquired,
	// no other read or writes are allowed
	defer m.Unlock()
	fmt.Println("writing", val)
	y = val
	time.Sleep(1 * time.Second)

}

func PrintY(wg *sync.WaitGroup, m *sync.RWMutex) {
	defer wg.Done()
	//no one can write when read lock is acquired,
	// there could be unlimited number of reads
	m.RLock() // Read Lock
	defer m.RUnlock()
	fmt.Println("read", y)
	time.Sleep(4 * time.Second)
}
