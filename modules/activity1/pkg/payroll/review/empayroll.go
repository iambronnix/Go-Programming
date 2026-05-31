package payroll

import (
	"fmt"
)

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}
type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}
type Manager struct {
	Individual     Employee
	Salary         float64
	CommissionRate float64
}
type Payer interface {
	pay() (string, float64)
}

func (d Developer) pay() (string, float64) {
	netPay := d.HourlyRate * d.HoursWorkedInYear
	return fmt.Sprintf("%v", d.Individual), netPay
}
func (m Manager) pay() (string, float64) {
	netPay := m.Salary + (m.CommissionRate * m.Salary)
	fullNames := m.Individual.FirstName + " " + m.Individual.LastName
	return fmt.Sprintf("%v", fullNames), netPay
}
