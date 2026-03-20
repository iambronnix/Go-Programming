package main

import (
	"fmt"
)

func main() {
	m := "Uncle Bob"
	greeting(m)
	fmt.Printf("Hi from main: %s\n", m)
}
func greeting(name string) {
	fmt.Printf("%s\n", name)
	s := "Slacker"
	fmt.Printf("Greeting %s\n", s)
}
