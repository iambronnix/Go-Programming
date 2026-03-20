package main

import (
	"errors"
	"fmt"
	"os"
)

func message() (message string, err error) {
	message = os.Args[1]
	if message == "hi" {

		err = fmt.Errorf("say bye\n") //err is being shadowed during return calling
		return "", err                //only accesible within the curly brackets

	}
	err = errors.New("Required string not passed\n")
	return "", err

}
func checkTerminal(message string) (string, error) {
	if os.Args[1] != message {
		return "", errors.New("No string was passed into the terminal")
	}
	return message, nil
}
func main() {
	_, f := checkTerminal(os.Args[1])
	_, m := message()

	fmt.Print(m)
	fmt.Println(f)
}
