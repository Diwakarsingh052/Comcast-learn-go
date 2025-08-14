package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wgTask := new(sync.WaitGroup)
	//wg.Add(5)

	for i := 1; i <= 3; i++ {
		wg.Add(1) // here the counter would be 5
		go func() {
			defer wg.Done() // defer must be inside the body of the goroutine
			fmt.Println("working on task id", i)

			wgTask.Add(1)
			go SubTask(i, wgTask)

			fmt.Println("working on task id, about to finish", i)
			wgTask.Wait()
		}()
	}

	wg.Wait()
}

func SubTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("doing some sub task", id)
}
