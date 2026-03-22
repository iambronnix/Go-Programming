package main

import (
	"fmt"
)

func main() {
	defer done()
	fmt.Println("Main: Start")
	fmt.Println("Main: End")
}
func done() {
	fmt.Println("Now I am done")
}
