package main

import "fmt"

func main() {

	counter := 4
	x := decrement(counter) //each call to x() runs the anonymous function
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}
func decrement(i int) func() int {
	return func() int {
		i-- //decrement by 1
		return i
	}
}
