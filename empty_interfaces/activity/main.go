package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}
type cat struct {
	name string
}

func main() {
	c := cat{name: "meozy"}
	i := 99
	b := false
	str := "testing"
	catDetails(c)
	emptyDetails(c)
	emptyDetails(i)
	emptyDetails(b)
	emptyDetails(str)
}
func (c cat) Speak() string {
	return "meoow meozyyy"
}
//Any type can be passed to the function since all types implement interface{}
implement the empty interface{}
func emptyDetails(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func catDetails(i Speaker) {
	fmt.Printf("(%v,%T)\n", i, i)
}
