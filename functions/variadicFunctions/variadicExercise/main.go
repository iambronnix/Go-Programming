package main

import "fmt"

func main() {
	i := []int{5, 10, 15}
	fmt.Println(sum(5, 4)) //calls the sum function by passing (5, 4) as args
	fmt.Println(sum(i...)) //calls the sum function passing the original args

}
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
