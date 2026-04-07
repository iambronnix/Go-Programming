package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

type person struct {
	name      string
	age       int
	isMarried bool
}

func main() {
	d := person{
		name:      "erick",
		age:       21,
		isMarried: true,
	}
	fmt.Println(d.String()) //calling the String method of the person struct and printing the result
	fmt.Println(d.Speak())  //calling the Speak method of the person struct and printing the result
}
func (p person) String() string { //String method returns a string representation of the person struct
	return "hey my name is :" + p.name
}
func (p person) Speak() string {
	return fmt.Sprintf("\nand i'm %v(and isMarried:%v)", p.age, p.isMarried) //Speak method returns a string representation of the person struct
}
