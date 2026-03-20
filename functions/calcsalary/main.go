package main

import (
	"fmt"
)

type bluePrint func(int int) int
func main(){
	a1 := salaryCalc(developer,2000,500)
}

func developer(baseSalary, monthlyBonus int) int {
	totalSalary := baseSalary + monthlyBonus
	return totalSalary
}
func manager(baseSalary, monthlyBonus int) int {
	totalSalary := baseSalary + monthlyBonus
	return totalSalary
func salaryCalc(baseSalary, monthlyBonus int, f bluePrint) {
	pay := f(baseSalary, monthlyBonus)
	fmt.Print(b(baseSalary, monthlyBonus))
}
