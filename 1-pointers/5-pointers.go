package main

import "fmt"

func main() {

	var p *int
	update(&p)
	fmt.Println(*p)
}
func update(p **int) {
	var x int = 10
	*p = &x
}
