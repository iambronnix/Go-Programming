package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourlyRate  = errors.New("invalid hourly rate")
	ErrHoursWorked = errors.New("invalid hours worked per week")
	ErrInvalidArgs = errors.New("invalid args")
)

func main() {

	pay := payDay(80, 50)
	fmt.Println(pay)
}
func payDay(hoursWorked, hourlyRate int) int {
	report := func() { // provides details of args provided inthe main()...provides insight why the program panics
		fmt.Printf("HoursWorked: %d\nHourlyRate: %d\n", hoursWorked, hourlyRate)
	}
	defer report()
	// if the data is invalid, the program panics and pass the argument of ErrHoulryRate or ErrHoursWorked
	if hourlyRate < 10 || hourlyRate > 75 {
		panic(ErrHourlyRate)
	}
	if hoursWorked < 0 || hoursWorked > 80 {
		panic(ErrHoursWorked)
	}
	// if panic occurs, the defer(), report(), will give the caller some insights into why the panic occured
	//panic will bubble up the stack to the main() function and execution will stop immediately
	if hoursWorked > 40 {
		hoursOver := hoursWorked - 40
		overTime := hoursOver * 2
		regularPay := hoursWorked * hourlyRate
		return regularPay + overTime
	}
	return hoursWorked * hourlyRate
}
