package main

import "fmt"

func main() {
	calculator(add, 5, 7)
	calculator(subtract, 9, 4)
}
func calculator(f func(int, int) int, i, j int) { //function type that takes two integers and returns an integer
	fmt.Println(f(i, j))
}
func add(i, j int) int { //function that matches the signature of the calculator function type
	return i + j
}
func subtract(i, j int) int { //function that matches the signature of the calculator function type
	return i - j
}
