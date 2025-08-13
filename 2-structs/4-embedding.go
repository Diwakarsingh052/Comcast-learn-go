package main

import (
	"fmt"
	"runtime"
)

type User struct {
	name  string
	email string
}

func (u *User) UpdateEmail(email string) {
	u.email = email
}

type BookAuthor struct {
	// anonymous field, field without a name
	// this case the field name would be taken from the type name
	//  if user is implementing an interface
	// than in embedding we get an automatic implementation of the interface
	User    // embedding // while embedding, we write the type name
	address string
}

func (b *BookAuthor) UpdateAddress(address string) {
	b.address = address
}
func (b *BookAuthor) PrintDetails() {
	fmt.Printf("%+v\n", b)
}

type MovieActor struct {
	u         User // not embedding
	movieList []string
}

func (m *MovieActor) AddMovie(movie string) {
	m.movieList = append(m.movieList, movie)
}
func (m *MovieActor) PrintDetails() {
	fmt.Printf("%+v\n", m)
}

func main() {
	fmt.Println(runtime.GOMAXPROCS(1))
	u := User{
		name:  "Rahul",
		email: "rahul@gmail.com",
	}

	b := BookAuthor{
		User:    u,
		address: "Bangalore",
	}
	m := MovieActor{
		u:         u,
		movieList: []string{"Inception", "The Matrix"},
	}

	b.UpdateEmail("rahul@example.com")
	b.PrintDetails()

	m.u.UpdateEmail("rahul11@gmail.com")
	m.AddMovie("The Matrix Reloaded")
	m.PrintDetails()
}
