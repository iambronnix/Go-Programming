package main

import (
	"fmt"
)

func main() {
	nums("Phyll", 99, 100)
}
func nums(str string, i ...int) {
	fmt.Println(str)
	fmt.Println(i)

}

//func f(parameterName ...Type)
