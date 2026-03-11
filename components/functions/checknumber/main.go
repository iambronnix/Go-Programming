package main

import (
	"fmt"
)

func checkNum(mynum int) {
	for i := 0; i <= mynum; i++ {
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
func main() {
	fmt.Println("Main is in control")
	checkNum(10)
	fmt.Println("Back to main")
}
