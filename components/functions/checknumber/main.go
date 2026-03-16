package main

import (
	"fmt"
)

func checkNum(i int) (int, string) {
	switch {
	case i%2 == 0:
		return i, "even"
	default:
		return i, "odd"

	}

}
func main() {
	for i := 0; i <= 15; i++ {

		_, result := checkNum(i) //i've used blank return to ignore the int return value...alternatively you can parse that int to a parameter and //
		fmt.Printf("Results:  %s\n", result)
	}
}
