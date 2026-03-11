package main

import (
	"fmt"
)

func checkNum(start, end int) {
	for i := start; i <= end; i++ {
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
func main() {
	fmt.Println("Main is in control")
	checkNum(10, 20)
	fmt.Println("Back to main")
}
