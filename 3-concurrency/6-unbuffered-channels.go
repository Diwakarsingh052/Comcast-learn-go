package main

func main() {
	// size is 0
	ch := make(chan int) // unbuffered channel
	// ch
	//go ->[]    <-go
	// receiver would subscribe to a channel, and sender would send the value
	// to receiver goroutine directly in case of unbuffered channel
}
