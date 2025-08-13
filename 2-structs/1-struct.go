package main

import "fmt"

func main() {
	type Person struct {
		name    string // fields of struct
		age     int
		address string
	}

	var p Person
	p1 := Person{
		name: "john",
		age:  30,
	}
	p.name = "John"

	fmt.Printf("%+v", p) // it can print field value pairs
	fmt.Printf("%#v", p1)
}
