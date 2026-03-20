package main

import "fmt"

func main() {
	v := square(9)
	fmt.Println(v())                 //calls the function returned by square(9), which calculates the square of 9 and returns 81
	fmt.Printf("Type of v: %T\n", v) //returns the type of v, which is func() int, indicating that v is a function that takes no arguments and returns an integer
}
func square(x int) func() int { //square is a function that takes an integer and returns a function that takes no arguments and returns an integer
	f := func() int {
		return x * x
	}
	return f
}
