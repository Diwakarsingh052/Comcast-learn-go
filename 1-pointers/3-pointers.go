package main

import "fmt"

func main() {
	x := 10 // x80
	p := &x
	updateVal(p) // x80

	fmt.Println(x)
	fmt.Println(*p)
}

func updateVal(px *int) { // px = x80

	var randomVal = 200 // x100
	// we have changed the reference of px to store a new variable
	px = &randomVal // px = x100
	// this would not update the main function x variable, it would update randomVal
	*px = 100
	fmt.Println(randomVal)

}
