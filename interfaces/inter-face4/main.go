package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}
type cat struct{}

func main() {
	d := cat{}
	chatter(d)
}
func (c cat) Speak() string { //Speak method returns a string
	return "Purr Meoww"
}
func chatter(s Speaker) { //chatter function takes a Speaker interface as an argument and calls the Speak method of the Speaker interface and prints the result
	fmt.Println(s.Speak())
}
