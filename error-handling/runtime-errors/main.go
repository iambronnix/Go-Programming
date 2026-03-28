package main

import "fmt"

func main() {
	nums := []int{2, 4, 6, 8}
	total := 0
	for i := 0; i <= 10; i++ { //This code will cause a runtime error because it tries to access an index of the nums slice that is out of range. The nums slice has a length of 4, so valid indices are 0 to 3. When i reaches 4, it will cause an "index out of range" error.
		total += nums[i]
	}
	fmt.Println("Total: ", total)
}
