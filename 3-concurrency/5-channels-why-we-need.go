package main

import "sync"

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	//a := go workV6(wg) // we can't return values from the goroutines
	// we need channel to communicate between goroutines
	wg.Wait()

}

func workV6(wg *sync.WaitGroup) int {
	defer wg.Done()
	return 10
}
