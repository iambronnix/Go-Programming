package main

import "fmt"

func main() {
	v := square(9)
	fmt.Println(v())
	fmt.Printf("Type of v: %T", v)
}
func square(x int) func() int {
	f := func() int {
		return x * x
	}
	return f
}
