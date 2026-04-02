package main

import (
	"errors"
	"fmt"
)

func main() {
	n := func() {
		fmt.Println("Defer in test") //going up the call stack, the next function is the main(), and it's deferred function, n() will execute
	}
	defer n()
	msg := "good-bye"
	message(msg)
}
func message(msg string) {
	f := func() {
		fmt.Println("Defer in message func") //deferred f(), runs in the message() function
	}
	defer f()
	if msg == "good-bye" { //when panic occurs,message() runs the defer statement within the panicking(), message()
		panic(errors.New("something went wrong"))
	}

}
