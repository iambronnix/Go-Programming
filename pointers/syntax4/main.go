package main

import (
	"fmt"
)

func main() {
	a, b := 5, 10
	swap(&a, &b)                                      //we need to pass the address of a and b to the function, so that it can modify the original variables.
	fmt.Println("successful swap\n", a == 10, b == 5) //after the swap function is called, a should be 10 and b should be 5.

}
func swap(a *int, b *int) { //a and b are pointers to int, so we need to dereference them to access the values they point to.
	*a = 10
	*b = 5

}
