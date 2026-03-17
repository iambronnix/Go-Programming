package main

import (
	"fmt"
)

func main() {
	increment := increment()
	fmt.Println(increment())
	fmt.Println(increment())
}
func increment() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
