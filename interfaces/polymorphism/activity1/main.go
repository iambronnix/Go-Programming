package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}
type cat struct{}
type dog struct{}
type person struct {
	name string
}

func main() {
	c := cat{}
	d := dog{}
	p := person{name: "erick"}
	catSpeak(c)
	dogSpeak(d)
	personSpeak(p)
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
func catSpeak(c cat) {
	fmt.Println(c.Speak())
}
func dogSpeak(d dog) {
	fmt.Println(d.Speak())
}
func personSpeak(p person) {
	fmt.Println(p.Speak())
}
