package main

import (
	"fmt"
)

func main() {
	age := 25
	name := "Meozy"
	defer personAge(name, age) //defer statement will execute after the main function finishes, but it will use the value of age at the time of defer statement execution, which is 25.
	age *= 2                   //double the age after the defer statement
	fmt.Printf("Age double %d.\n", age)
}
func personAge(name string, i int) { //this function will be called after the main function finishes, but it will use the value of age at the time of defer statement execution, which is 25.
	fmt.Printf("%s is %d.\n", name, i)
}
