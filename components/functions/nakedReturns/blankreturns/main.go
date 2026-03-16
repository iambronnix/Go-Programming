package main

import (
	"fmt"
)

func message() (message string, err error) {
	message = "hi"
	if message == "hi" {
		err := fmt.Errorf("say bye\n") //err is being shadowed during return calling
		return
	}
	return

}
