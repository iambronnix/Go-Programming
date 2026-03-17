package main

import "fmt"

func main() {
	m := func() int {
		return sum(2, 5, 8)

	}()
	a := sum(m)
	fmt.Printf("The sum is: %v\nThe values are:%v\n", a, m)

}
func sum(m ...int) int {
	init := 0
	for _, value := range m {
		init += value
	}
	return init
}
