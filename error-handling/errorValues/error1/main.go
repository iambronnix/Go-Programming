package main

import (
	"errors"
	"fmt"
)

func main() {
	ErrBadData := errors.New("Some bad data\n") //New() takes a string as an argument, converts it into  *errors.errorString...returns an error value
	fmt.Printf("ErrBadData type: %T\n", ErrBadData)
}
