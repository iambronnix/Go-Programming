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
		num, result := checkNum(i)
		fmt.Printf("Results: %d %s\n", num, result)
	}
}
