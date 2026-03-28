package main

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}
type Details struct {
	Individual Employee
	Payrate    int
	WorkWeek   [7]int
}
type WeekDay int

const (
	Sunday WeekDay = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
