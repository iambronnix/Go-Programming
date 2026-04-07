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
	fmt.Println(d.Speak())
	fmt.Println(d)

}
func (c cat) Speak() string {
	return "Purr Meow"
}
func (c cat) String() string {
	return fmt.Sprintf("%v (%v years old)\n", c.name, c.age)
}
