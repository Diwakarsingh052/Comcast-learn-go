package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := openFileV2("test.txt")
	if err != nil {
		log.Println(err, "in main")
		return
	}
	info, _ := f.Stat()
	fmt.Println(info.Name())

}

func openFileV2(fileName string) (*os.File, error) {
	f, err := os.Open(fileName)
	if err != nil {
		//errors.Is can check if an error was wrapped inside the chain or not
		//  if an error was found in the chain, you now know what exactly went wrong
		// you might want to take some actions to fix the issue
		//or maybe just log the additional details
		if errors.Is(err, os.ErrNotExist) {
			log.Println("attempting to create file")
			f, err := os.Create(fileName)
			if err != nil {
				return nil, err
			}
			return f, nil

		}

		// if error happened because of another reason other than ErrNotExist
		return nil, err
	}

	return f, nil
}
