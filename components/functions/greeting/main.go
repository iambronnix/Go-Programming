package main

import (
	"fmt"
)

func main() {
	greeting(6, "Meozy")
}
func greeting(age int, name string) {
	fmt.Printf("%s is %v months\n", name, age)
}
