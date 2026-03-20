package main

import "fmt"

type calc func(int, int) string //calc is a function type that takes two integers and returns a string

func main() {
	calculator(add, 5, 7)
	calculator(sub, 9, 4)
}
func add(i, j int) string { //function that matches the signature of the calc type
	result := i + j
	return fmt.Sprintf("Added %d + %d = %d", i, j, result)
}
func sub(i, j int) string { //function that matches the signature of the calc type
	result := i - j
	return fmt.Sprintf("Subtracted %d - %d = %d", i, j, result)
}
func calculator(f calc, i, j int) {
	fmt.Println(f(i, j)) //delegates the operation to the function passed as an argument
}
