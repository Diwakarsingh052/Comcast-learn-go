package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	//wg.Add(5)

	for i := 1; i <= 3; i++ {
		wg.Add(1) // here the counter would be 5
		go func() {
			wgTask := new(sync.WaitGroup)
			defer wg.Done() // defer must be inside the body of the goroutine
			fmt.Println("working on task id", i)

			// using another waitgroup to keep track of subtasks we are running
			wgTask.Add(1)
			go SubTask(i, wgTask)

			fmt.Println("working on task id, about to finish", i)
			// waiting for subtask to finish before we quit the parent goroutine
			wgTask.Wait()
		}()
	}

	wg.Wait()
}

func SubTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("doing some sub task", id)
}
