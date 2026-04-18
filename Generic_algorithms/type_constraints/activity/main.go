package main

import (
	"errors"
	"fmt"
)

type Number interface {
	int | float64 //types that will be allowed
}

var ErrArray = errors.New("an empty array is being passed")

func main() {
	maxInt, _ := findMax([]int{1, 32, 5, 8, 10, 11})
	maxFloat, _ := findMax([]float64{1.1, 32.1, 5.1, 8.1, 10.1, 11.1})
	fmt.Printf("maximum integer:%d\tmaximum float:%v\n", maxInt, maxFloat)
}
func findMax[Num Number](num []Num) (Num, error) {
	if len(num) == 0 {
		panic(ErrArray)
	}
	max := num[0]
	for _, key := range num {
		if key > max {
			max = key
		}
	}
	return max, nil
}
