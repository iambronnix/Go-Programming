package main

import (
	"errors"
	"fmt"
)

var ErrArray = errors.New("you're passing an empty array ")

func main() {
	max, _ := findMaxInt([]int{34, 32, 56, 87, 12, 34, 22})
	max2, _ := findMaxFloat([]float64{63, 23, 34, 54, 67, 31})
	fmt.Printf("max float value:%v\nmax int value:%v\n", max2, max)
}
func findMaxInt(nums []int) (int, error) {
	if len(nums) == 0 {
		panic(ErrArray)
	}
	max := nums[0]
	for _, key := range nums {
		if max > key {
			max = key
		}
	}
	return max, ErrArray
}
func findMaxFloat(nums []float64) (float64, error) {
	if len(nums) == 0 {
		panic(ErrArray)

	}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max, ErrArray
}
