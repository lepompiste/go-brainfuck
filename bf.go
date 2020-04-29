package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ExecutionError is error of execution datatype structure
type ExecutionError struct {
	Message string
	ID      int
}

func check(e error) {
	if e != nil {
		fmt.Println("Une erreur est survenue...")
		os.Exit(-1)
	}
}

// ExitWithError quits the program printing message
func ExitWithError(error ExecutionError) {
	os.Exit(error.ID)
	fmt.Println(error.Message)
}

// TODO manual reading
// TODO add only instructions to global instructions slice
func main() {
	if len(os.Args) == 2 {
		data, err := ioutil.ReadFile(os.Args[1]) // data []byte, err error
		check(err)
		instructions = data
		Execute()
	} else {
		fmt.Println("Usage : go-bf <file>")
	}
}
