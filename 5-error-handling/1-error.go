package main

import (
	"errors"
	"fmt"
	"log"
	"runtime/debug"
)

var user = make(map[int]string)

// variable which would be used to store errors should start with word Err
var ErrRecordNotFound = errors.New("record not found")

func main() {
	//var err1 error // default value of err is nil
	// avoid if else branching while checking for error, use return to stop the function
	name, err := FetchRecord(10) // use _ to ignore a return value
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(name)
	name, err = FetchRecord(100)
	if err != nil {
		return
	}
	name, err = FetchRecord(200)
	if err != nil {
		return

	}
}

// error must be the last value to be returned from function

func FetchRecord(id int) (string, error) {
	name, ok := user[id]
	if !ok {
		fmt.Println(string(debug.Stack()))
		// whenever error happens, set other values to their defaults
		return "", ErrRecordNotFound
	}
	// default value of error is nil // which means no error happened
	return name, nil
}
