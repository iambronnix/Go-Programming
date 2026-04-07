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
func (c cat) Speak() string {
	return "Purr Meoww"
}
func chatter(s Speaker) {
	fmt.Println(s.Speak())
}
