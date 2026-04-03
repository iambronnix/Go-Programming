package main

import (
	"errors"
	"fmt"
)

var (
	ErrHoulryRate  = errors.New("Invalid hourly rate")
	ErrHoursWorked = errors.New("Invalid hours worked per week ")
)

func main() {
	//call payDay() with various args
	pay := payDay(100, 25)
	fmt.Println(pay)
	pay = payDay(100, 200)
	fmt.Println(pay)
	pay = payDay(60, 25)
	fmt.Println(pay)
}

func payDay(hoursWorked, hourlyRate int) int {
	defer func() {
		if r := recover(); r != nil { // check for the return value from recovery() func...if r isn't nil, means a panic occurs
			if r == ErrHoulryRate { //evaluate r and see whether it equals one of the error values- ErrHourlyRate or ErrHoursWorked
				fmt.Printf("hourly rate: %d\nerr: %v\n\n", hourlyRate, r)
			}
			if r == ErrHoursWorked {
				fmt.Printf("hours worked: %d\nerr: %v\n\n", hoursWorked, r)
			}

		}
		fmt.Printf("Pay was calculated based on:\nhours worked: %d\nhourly Rate: %d\n", hoursWorked, hourlyRate)
	}()
	if hourlyRate < 10 || hourlyRate > 75 {
		panic(ErrHoulryRate)
	}
	if hoursWorked < 0 || hoursWorked > 80 {
		panic(ErrHoursWorked)
	}
	if hoursWorked > 40 {
		hoursOver := hoursWorked - 40
		overTime := hoursOver * 2
		regularPay := hoursWorked * hourlyRate
		return regularPay + overTime
	}
	return hoursWorked * hourlyRate
}
