package main

import (
	"flag"
	"fmt"
)

func main() {
	var j int
	flag.IntVar(&j, "myvalue", 0, "use myvalue flag")
	flag.Parse()
	x := func(i int) int { //anonymous function assigned to x
		return i * i
	}
	fmt.Printf("The square of %d is %d\n", j, x(j))
}
