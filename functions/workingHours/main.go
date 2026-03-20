package main

import (
	"fmt"
)

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}
type Developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek   [7]int
}
type WeekDay int

const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	d := Developer{Individual: Employee{Id: 1, FirstName: "Jayden", LastName: "Mutiri"}, HourlyRate: 10}
	d.LogHours(Monday, 8)
	d.LogHours(Tuesday, 10)
	fmt.Println("Hours worked on Monday: ", d.WorkWeek[Monday])
	fmt.Println("Hours worked on Tuesday: ", d.Workweek[Tuesday])
	fmt.Printf("Hours worked this week: %d", d.HoursWorked())
}
func (d *Developer) LogHours(day WeekDay, hours int) {
	d.WorkWeek[day] = hours
}
func (d *Developer) HoursWorked() int {
	total := 0
	for _, v := range d.WorkWeek {
		total += v
	}
	return total
}
