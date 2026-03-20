package main

import (
	"errors"
	"fmt"
)

func message() (message string, err error) {
	message = "james"
	if message == "hi" {
		err = fmt.Errorf("say bye\n") //err is being shadowed during return calling
		return "", err                 //only accesible within the curly brackets

	}
	err = errors.New("another string was passed")
	return "", err

}
func main() {
	_, m := message()

	fmt.Print(m)
}

	fmt.Print(m)
	fmt.Println(f)
}
