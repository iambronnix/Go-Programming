package main

import (
	"fmt"
)

func main() {
	//In each instance of the variable declaration, each variable is declared as an empty interface
	var str interface{} = "some string"
	var i interface{} = 42
	var b interface{} = true
	fmt.Println(str)
	fmt.Println(i)
	fmt.Println(b)
}
