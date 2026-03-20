package main

import (
	"fmt"
)

func main() {
	nums(99, 100)

}
func nums(i ...int) {
	fmt.Println(i)
	fmt.Printf("%T\n", i)
	fmt.Printf("Len: %d\n", len(i))
	fmt.Printf("Cap: %d\n", cap(i))
}
