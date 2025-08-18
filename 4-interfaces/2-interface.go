package main

import "fmt"

// Polymorphism means that a piece of code changes its behavior depending on the
// concrete data it’s operating on // Tom Kurtz, Basic inventor

type Reader interface {

	// interfaces are automatically implemented if method signature is same as of interfaces
	// only method signatures could be added to an interface
	Read(b []byte) (int, error)
}
type File struct {
	Name string
}

func (f File) Read(b []byte) (int, error) {
	fmt.Println("reading files and processing them", f.Name)
	return 0, nil
}

type IO struct {
	name string
}

func (i IO) Read(b []byte) (int, error) {
	fmt.Println("reading and processing io ", i.name)
	return 0, nil
}

func DoSomeReading(f File, i IO) {
	f.Read(nil)
	i.Read(nil)
}

// r can accept any kind of type which implements the Reader interface

func DoSomeReadingV2(r Reader) {
	//piece of code changes its behavior depending on the
	//// concrete data it’s operating on

	// if file is passed Read would be called from the file struct,
	//or if Io is passed Read would be called from IO
	r.Read(nil)
	fmt.Printf("type of r %T\n", r)
}
func main() {
	f := File{Name: "test.txt"}
	i := IO{name: "IO"}
	DoSomeReadingV2(f)
	DoSomeReadingV2(i)
}
