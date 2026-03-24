package main

import (
	"fmt"
)

func main() {
	age := 25
	name := "Meozy"
	defer personAge(name, age)
	age *= 2
	fmt.Printf("Age double %d.\n", age)
}
func personAge(name string, i int) {
	fmt.Printf("%s is %d.\n", name, i)
}
