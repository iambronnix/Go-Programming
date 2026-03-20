package main

import "fmt"

type calc func(int, int) string

func main() {
	calculator(add, 5, 7)
	calculator(sub, 9, 4)
}
func add(i, j int) string {
	result := i + j
	return fmt.Sprintf("Added %d + %d = %d", i, j, result)
}
func sub(i, j int) string {
	result := i - j
	return fmt.Sprintf("Subtracted %d - %d = %d", i, j, result)
}
func calculator(f calc, i, j int) {
	fmt.Println(f(i, j))
}
