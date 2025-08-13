package main

import (
	"fmt"
)

func main() {
	// assume p address is x80 and value is nil
	var p *int           // nil
	updateNilPointer(&p) // passing address of p (x80)

	// pointer would have updated value because of double pointer
	fmt.Println(*p)

	//var user *user // nil
	//errors.As()
	//json.Unmarshal(&user) // double pointer is must,
	//without it json.Unmarshal can't change the values of pointer

}

func updateNilPointer(p1 **int) {
	x := 10
	// trying to access the value of p1
	// which is also another pointer named as p from the main function
	*p1 = &x

}
