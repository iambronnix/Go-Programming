package main

import (
	"errors"
	"fmt"
)

var Errpass = errors.New("invalid input being passed\n")

func main() {
	minInt, _ := findMinVal([]int{1, 32, 5, 8, 10, 11})
	minFloat, _ := findMinVal([]float64{1.1, 32.1, 5.1, 8.1, 10.1, 11.1})
	fmt.Printf("Minimum integer:%v\tMinimum float:%v\n", minInt, minFloat)
}

func findMinVal[Num int | float64](num []Num) (Num, error) {
	if len(num) == 0 {
		panic(Errpass)
	}
	min := num[0]
	for _, key := range num {
		if key < min {
			min = key
		}
	}
	return min, nil
}
