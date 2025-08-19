package main

import (
	"fmt"
	"log"
	"runtime/debug"
)

// panic is a runtime exception
// we need to decide at what leve we need to stop panic
// the level we stop panic must stop or that function must stop that is recovering the panic

// if the caller function doesn't depends on the called function you can stop panic propagation back
// by calling the recovery function in defer in the function that is getting called
func main() {

	DoSomething()
	fmt.Println("doing more stuff in main, which is not related to DoSomething")
}

func DoSomething() {
	// we need to decide where we would recover from the panic
	// the func where we decide to recover the panic needs to stop

	// RecoverPanic would recover the current function from panic, but the function needs to stop
	// it can't continue executing

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
