package main

import (
	"errors"
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

func main() {

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
func payDetails(p Payer) string {
	fullNames, netPay := p.pay()
	return fmt.Sprintf("Employer:%v\tNet pay\t:%v", fullNames, netPay)
}
func convertComment(str string, err error) (int, error) {
	err = errors.New("Invalid comment")
	switch str {
	case "Excellent":
		return 5, nil
	case "Good":
		return 4, nil
	case "Poor":
		return 2, nil
	case "Unsatisfactory":
		return 1, nil
	default:
		panic(err)

	}
}
