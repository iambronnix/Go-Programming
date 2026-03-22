package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("Finally called the anonymous() func")
	}()
	fmt.Println("Starting the main() func")
	fmt.Println("Ending execution of the main() func")
}
