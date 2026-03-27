package main

import "fmt"

type Developer struct { //Developer struct has two fields, one is of type Employee and another is of type int and array of int.
	Individual Employee
	HourlyRate int
	WorkWeek   [7]int
}
type Employee struct {
	Id        int
	FirstName string
	LastName  string
}
type Weekday int //Weekday is a custom type that represents the days of the week. It is defined as an integer type, and the constants for each day of the week are defined using iota, which automatically assigns successive integer values starting from 0 for Sunday, 1 for Monday, and so on.

const (
	Sunday Weekday = iota //iota is a predeclared identifier that represents successive untyped integer constants. It resets to 0 whenever the word const appears in the source and increments by 1 after each const specification in the same block.
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	d := Developer{Individual: Employee{Id: 1, FirstName: "Erick", LastName: "Ndeto"}, HourlyRate: 10}
	x := nonLoggedHours()
	fmt.Println("Tracking hours worked thus far today :", x(2))
	fmt.Println("Tracking hours worked thus far today :", x(3))
	fmt.Println("Tracking hours worked thus far today :", x(5))
	fmt.Println()
	d.LogHours(Monday, 8)
	d.LogHours(Tuesday, 10)
	d.LogHours(Wednesday, 12)
	d.LogHours(Thursday, 21)
	d.LogHours(Friday, 6)
	d.LogHours(Saturday, 5)
	d.PayDetails()

}
func (d *Developer) LogHours(day Weekday, hours int) {
	d.WorkWeek[day] = hours //LogHours is a method that takes a pointer receiver of type Developer, a parameter of type Weekday, and a parameter of type int. It assigns the value of hours to the corresponding index in the WorkWeek array based on the value of day.
}
func (d *Developer) HoursWorked() int { //HoursWorked is a method that takes a pointer receiver of type Developer and returns the total number of hours worked in a week by summing up the values in the WorkWeek array.
	total := 0
	for _, v := range d.WorkWeek {
		total += v
	}
	return total
}
func (d *Developer) PayDay() (int, bool) {
	if d.HoursWorked() > 40 {
		hoursOver := d.HoursWorked() - 40
		overTime := hoursOver * 2 * d.HourlyRate
		regularPay := 40 * d.HourlyRate
		return regularPay + overTime, true
	}
	return d.HoursWorked() * d.HourlyRate, false
}
func nonLoggedHours() func(int) int {
	total := 0
	return func(i int) int { //nonLoggedHours is a function that returns a closure. The closure takes an integer parameter and adds it to the total variable, which is defined in the outer function. Each time the closure is called, it updates the total and returns the new total.
		total += 1
		return total
	}
}
func (d *Developer) PayDetails() {
	for i, v := range d.WorkWeek {
		switch i {
		case 0:
			fmt.Println("Sunday hours: ", v)
		case 1:
			fmt.Println("Monday hours:", v)
		case 2:
			fmt.Println("Tuesday hours:", v)
		case 3:
			fmt.Println("Wednesday hours:", v)
		case 4:
			fmt.Println("Thursday hours:", v)
		case 5:
			fmt.Println("Friday hours:", v)
		case 6:
			fmt.Println("Saturday hours:", v)
		}
	}
	fmt.Printf("\nHours Worked this week: %d\n", d.HoursWorked())
	pay, overtime := d.PayDay()
	fmt.Println("Pay for the week: $", pay)
	fmt.Println("Is this overtime pay: ", overtime)
	fmt.Println()
}
