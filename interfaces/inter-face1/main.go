package main

import (
	"fmt"
)

type Speaker interface { //Speaker interface has a method Speak() that returns a string
	Speak() string
}
type cat struct { //cat struct implements the Speaker interface
}

func main() {
	c := cat{}             //creating an instance of cat struct
	fmt.Println(c.Speak()) //calling the Speak method of the cat struct and printing the result
	c.Greeting()           //calling the Greeting method of the cat struct
}
func (c cat) Speak() string { //Speak method returns a string
	return "Purr Meow"
}
func (c cat) Greeting() { //Greeting method prints a message
	fmt.Println("Meow, Meow!!!!mmmmmeeeeeeooooooowwwww")
}
