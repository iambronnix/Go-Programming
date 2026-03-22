package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("I was declared first.")
	}()
	defer func() {
		fmt.Println("I was declared second")
	}()
	defer func() {
		fmt.Println("I was declared third")
	}()
	f1 := func() {
		fmt.Println("Main: Start")
	}
	f2 := func() {
		fmt.Println("Main:End")
	}
	f1()
	f2()

}
