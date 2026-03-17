package main

import (
	"fmt"
)

func main() {
	i := 0
	incrementor := func() int {
		i += 1 //increments i by 1
		return i
	}
	fmt.Println(incrementor()) //first prints 1
	fmt.Println(incrementor()) // second time prints 2
	i += 10                    //this increases the value of i by 10 ...i can be accessed and changed outside the anonymous function
	fmt.Println(incrementor()) // prints 12
}
