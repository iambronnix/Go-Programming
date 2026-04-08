package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}
type cat struct{}
type person struct {
	name string
}
type dog struct{}

func saySomething(say ...Speaker) {
	for _, s := range say {
		fmt.Println(s.Speak())
	}
}
func main() {
	c := cat{}
	d := dog{}
	p := person{name: "erick"}
	saySomething(c, d, p)
}

func (c cat) Speak() string {
	return "Purr Meow"
}
func (d dog) Speak() string {
	return "woof woof"
}
func (p person) Speak() string {
	return "Hi my name is" + " " + p.name + "."
}
