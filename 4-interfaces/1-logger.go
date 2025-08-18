package main

import (
	"log"
	"os"
)

func main() {

	log.New(&os.File{}, "", 0)
	log.New(os.Stderr, "", 0)
	log.New(os.Stdout, "", 0)
}
