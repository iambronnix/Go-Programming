package main

import "fmt"

var name = "Gopher"

func init() {
	fmt.Println("Hello,", name)
}
func init() {
	fmt.Println("Second")
}
func init() {
	fmt.Println("Third")
}
func main() {
	fmt.Println("Hello, main function")
}
