package main

import (
	"app/internal/models"
	"app/internal/stores"
	"app/internal/stores/mysql"
	"app/internal/stores/postgres"
	"fmt"
)

func main() {
	p := postgres.NewConn()
	m := mysql.NewConn()
	fmt.Println()
	//var s stores.UserRepository
	// s = p
	// s.Create()
	u := models.User{
		Id:   1,
		Name: "Dev",
	}
	// using interface we would switch the implementation according to the concrete implementation
	stores.CreateUser(&p, u)

	stores.CreateUser(&m, u)

}
