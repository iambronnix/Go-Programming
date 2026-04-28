package main

import "fmt"

var name = "Gopher"

func init() {
	fmt.Println("hello", name)
}
func main() {
	fmt.Println("Hello main function")
}
