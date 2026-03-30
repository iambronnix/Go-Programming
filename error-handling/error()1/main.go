package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := "10"
	if s, err := strconv.Atoi(v); err == nil { //strconv.Atoi() returns int and an error
		fmt.Printf("%T, %v\n", s, s)

	} else {
		fmt.Println(err)
	}
	v = "s2"
	s, err := strconv.Atoi(v)
	if err != nil {
		fmt.Println(s, err)
	}
}
