package main

import "fmt"

func main() {
	a := 10

	p := &a
	fmt.Println("address of a", &a)
	fmt.Println("value of p", p)
	*p = 20 // dereference of pointer //  *// value at the address
	fmt.Println(a)
	update(&a)
	//update(p) // same as above
	fmt.Println(a)
	fmt.Println(*p)
}

func update(i *int) {
	fmt.Println("value of i", i)
	*i = 10
}
