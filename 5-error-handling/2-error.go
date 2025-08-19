package main

import (
	"log"
	"os"
)

func openFile(filename string) (*os.File, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func main() {
	f, err := openFile("test1.txt")
	if err != nil {
		log.Println(err, "some msg")
	}
	f.Close()

}

//[]
