package main

import (
	"fmt"
)

func main() {
	increment := increment() //assigns a variable to func() int that gets returned
	fmt.Println(increment())
	fmt.Println(increment())
}
func increment() func() int { // increment() is of func() int type
	i := 0              //only increment() has access to var i
	return func() int { //return type only increments i...note each call runs the anonymous function code
		i += 1
		return i
	}
}
