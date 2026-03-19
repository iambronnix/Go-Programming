package main

import (
	"fmt"
)

type calc func(int, int) string

func main() {
	calculator(add, 5, 6)
}
func add(i, j int) string {
	result := i + j
	return fmt.Sprintf("Added %d + %d = %d", i, j, result)
}
func calculator(f calc, i, j int) {
	fmt.Println(f(i, j))
}
