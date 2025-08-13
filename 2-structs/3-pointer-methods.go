package main

import (
	"fmt"
)

type author struct {
	name string
	age  int
}

func (a *author) UpdateName(name string) {

	a.name = name
	fmt.Printf("inside the method %+v\n", a)

}

func main() {
	//a := author{
	//	name: "John",
	//	age:  30,
	//}
	var a author
	a.UpdateName("Jim") // a address is copied to the receiver
	fmt.Println(a)
}
func nilPointer() {
	var a *author       // nil
	a.UpdateName("Jim") // panic // a is a nil pointer
	// updateName method is not a double pointer so that can't update a directly
	fmt.Println(a)

}
