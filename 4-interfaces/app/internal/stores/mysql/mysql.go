package mysql

import (
	"app/internal/models"
	"fmt"
	"sync"
)

// var UserDb map[int]models.User
//var conn *sql.DB // avoid setting connections in global variables
// anyone can change it and it is changed for all

//type Conn struct {
//	db *sql.DB
//}

type Conn struct {
	m      sync.Mutex
	userDb map[int]models.User
}

func NewConn() Conn {

	c := Conn{
		userDb: make(map[int]models.User, 100),
	}
	return c

}

// Creating a map to act as a data store

func (c *Conn) Create(u models.User) (models.User, bool) {
	c.m.Lock()
	defer c.m.Unlock()
	fmt.Println("Creating a user in mysql ", " u : ", u)
	//Need to check if user exists if yes then throw error else save
	c.userDb[u.Id] = u

	return u, true

}

func (c *Conn) Update(id int, name string) (models.User, bool) {
	c.m.Lock()
	defer c.m.Unlock()
	u, ok := c.userDb[id]

	if !ok {
		fmt.Println("User with id ", id, "Is not found for update")
		return models.User{}, false
	}
	fmt.Println("Updating a user in ", " u : ", u)

	u.Name = name
	return u, true
}
