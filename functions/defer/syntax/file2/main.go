package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("There goes the anonymous function")
	}()
	fmt.Println("Main: Start")
	fmt.Println("Main: End")

}
