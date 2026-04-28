package main

import "fmt"

var budgetCategories = make(map[int]string)

func init() {
	//data initialization and loading before the main() function executes
	fmt.Println("Initializing our budgetCategories")
	budgetCategories[1] = "Car insurance"
	budgetCategories[2] = "Mortage"
	budgetCategories[3] = "Electricity"
	budgetCategories[4] = "Retirement"
	budgetCategories[5] = "Vacation"
	budgetCategories[6] = "Groceries"
	budgetCategories[7] = "Car Payment"
}

func main() {
	for k, j := range budgetCategories {
		fmt.Printf("key:%d, value:%s\n", k, j)
	}
}
