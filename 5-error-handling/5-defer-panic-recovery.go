package main

import (
	"fmt"
	"log"
	"runtime/debug"
)

func main() {

	DoSomething()
	fmt.Println("doing more stuff in main, which is not related to DoSomething")
}

func DoSomething() {
	defer recoverPanic()
	UpdateSlice()
	fmt.Println("doing more things in DoSomething which is related to UpdateSlice")
}

func UpdateSlice() {

	var s []int
	s[0] = 100
	fmt.Println("Update Slice finished")
}

func recoverPanic() {
	// The built-in `recover` function can stop the process of panicking,
	//if it is called within a deferred function.

	// msg would have the actual panic message if that happened
	msg := recover()

	if msg != nil {
		// if this condition is true then panic happened
		log.Println(msg)
		fmt.Println()
		log.Println(string(debug.Stack()))

	}
}
