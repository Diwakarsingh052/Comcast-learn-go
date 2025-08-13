package main

import (
	"fmt"
)

var a int = 10

func main() {
	var p *int // nil
	fmt.Println(&p)
	// after calling the update value p is still nil, as we never updated the pointer
	updateValue(p)
	fmt.Println(*p) // panic // nil pointer dereference

}

func updateValue(x *int) {
	fmt.Println(&x)
	x = &a
	fmt.Println(*x)

}
