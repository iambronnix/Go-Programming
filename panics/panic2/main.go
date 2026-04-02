package main

import (
	"errors"
	"fmt"
)

func main() {
	msg := "good-bye"
	message(msg)
	fmt.Println("This line will not get printed") //if we were to print this line we would use defer before message()....that's while keeping the panic()
}
func message(msg string) { //func panics because the argument to the func message is "good-bye"
	if msg == "good-bye" {
		panic(errors.New("something went wrong")) //prints the error message.
	} //
}
