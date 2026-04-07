package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}
type cat struct {
	name string
	age  int
}

func main() {
	d := cat{
		name: "meozy",
		age:  3,
	}
	fmt.Println(d.Speak()) //calling the Speak method of the cat struct and printing the result
	fmt.Println(d)         //calling the String method of the cat struct and printing the result

}
func (c cat) Speak() string { //Speak method returns a string
	return "Purr Meow"
}
func (c cat) String() string { //String method returns a string representation of the cat struct
	return fmt.Sprintf("%v (%v years old)\n", c.name, c.age)
}
