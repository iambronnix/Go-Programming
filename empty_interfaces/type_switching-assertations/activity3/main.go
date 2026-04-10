package main

import (
	"fmt"
)

type cat struct {
	name string
}

func main() {
	c := cat{name: "meozy"}
	i := []interface{}{21, "This is how i code", true, c}
	typeExample(i)
}
func typeExample(i []interface{}) {
	for _, x := range i {
		switch v := x.(type) { //The for loop ranges over the slice of interfaces
		case int:
			fmt.Printf("%v is an int\n", v)
		case string:
			fmt.Printf("%v is a string\n", v)
		case bool:
			fmt.Printf("a bool %v\n", v)
		default:
			fmt.Printf("Unknown type %T\n", v)
		}
	}
}
