package main

import (
	"errors"
	"fmt"
)

var ErrArray = errors.New("you're passing an empty array ")

func main() {
	max, _ := findMaxGeneric([]int{34, 32, 56, 87, 12, 34, 22})
	max2, _ := findMaxGeneric([]float64{63, 23, 34, 54, 67, 31})
	fmt.Printf("max float value:%v\nmax int value:%v\n", max2, max)
}
func findMaxGeneric[Num int | float64](num []Num) (Num, error) {
	if len(num) == 0 {
		panic(ErrArray)
	}
	max := num[0]
	for _, key := range num {
		if max > key {
			max = key
		}

	}
	return max, nil

}
