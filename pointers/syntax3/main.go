package main

import (
	"fmt"
)

func add5Value(count int) { //count is a copy of the value passed to the function, so any changes made to count inside the function will not affect the original variable outside the function.
	count += 5
	fmt.Println("add5Value :", count)
}
func add5Point(count *int) { //count is a pointer to an int, so we need to dereference it to access the value it points to.
	*count += 5
	fmt.Println("add5point  :", *count)
}
func main() {
	var count int
	add5Value(count)
	fmt.Println("add5Value post:", count)
	add5Point(&count) //we need to pass the address of count to the function, so that it can modify the original variable.
	fmt.Println("add5point post:", count)
}
