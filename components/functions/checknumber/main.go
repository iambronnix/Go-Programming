package main

import (
	"fmt"
)

func checkNum() {
	for i := 0; i <= 30; i++ {
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
func main() {
	fmt.Println("Main is in control")
	checkNum()
	fmt.Println("Back to main")
}
