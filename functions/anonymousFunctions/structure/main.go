package main

import (
	"fmt"
)

func main() {
	message := "hey buddy\n"
	func(str string) { //anonymous func creation and parameter passing
		fmt.Print(str)
	}(message) //anonymous func execution point
}
