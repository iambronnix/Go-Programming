package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourlyRate  = errors.New("invalid hourly rate")
	ErrHoursWorked = errors.New("invalid hours worked per week")
)

func main() {
	pay, err := payDay(81, 50) //payDay() returns two values, an int and an error. We can assign those to pay and err variables
	if err != nil {
		fmt.Println(err)
	}
	pay, err = payDay(80, 75) //we can reuse the same variables to call payDay() again with different arguments
	if err != nil {
		fmt.Println(err)
	}
	pay, err = payDay(75, 50) //we can reuse the same variables to call payDay() again with different arguments
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pay)
}
func payDay(hoursWorked, hourlyRate int) (int, error) {
	if hourlyRate < 10 || hourlyRate > 75 {
		return 0, ErrHourlyRate //we return 0 for the int value and ErrHourlyRate for the error value
	}
	if hoursWorked > 40 { //if hoursWorked is greater than 40, we calculate the overtime pay
		hoursOver := hoursWorked - 40
		overTime := hoursOver * 2
		regularPay := hoursWorked * hourlyRate
		return regularPay + overTime, nil //we return the total pay and nil for the error value
	}
	return hoursWorked * hourlyRate, nil //if hoursWorked is less than or equal to 40, we calculate the regular pay and return it with a nil error
}
