package main

import (
	"errors"
	"fmt"
)

func main() {
	max, _ := findMaxInt([]int{34, 32, 56, 87, 12, 34, 22})
	fmt.Printf("max interger value:%v\n", max)
}
func findMaxInt(nums []int) (int, error) {
	if len(nums) == 0 {
		err := errors.New("an empty array")
		panic(err)

	}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max, nil
}
