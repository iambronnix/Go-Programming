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

		_, result := checkNum(i)
		fmt.Printf("Results:  %s\n", result)
	}
}
