package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	for i := 0; i <= 10; i++ { //this will cause a panic because we are trying to access an index that is out of range of the nums slice
		fmt.Println(nums[i])
	}
}
