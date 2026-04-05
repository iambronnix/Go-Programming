package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidRoutingNumber = errors.New("invalid routing number")
	ErrInvalidLastName      = errors.New("invalid last name")
)

type directDeposit struct {
	firstName        string
	lastName         string
	bankName         string
	routingNumber    int
	accountingNumber int
}

func main() {
	d := directDeposit{firstName: "Erick", lastName: "Ndeto", bankName: "NCBA", routingNumber: 35, accountingNumber: 228993988}
	defer d.validateLastName()
	if err := ErrInvalidLastName; err != nil {
		d.report(err)
	}
	defer d.validateRoutingNumber()
	if err := ErrInvalidRoutingNumber; err != nil {
		d.report(err)
		fmt.Printf("%v\n%v\n%v\n%v\n%v", d.firstName, d.lastName, d.bankName, d.routingNumber, d.accountingNumber)
	}

}

func (d *directDeposit) validateRoutingNumber() error {
	if d.routingNumber < 100 {
		return ErrInvalidRoutingNumber
	}
	return nil
}
func (d *directDeposit) validateLastName() error {
	if d.lastName == "" {
		return ErrInvalidLastName
	}
	return nil
}
func (d *directDeposit) report(error) {
	if err := ErrInvalidLastName; err != nil {
		panic(err)

	}
	if err := ErrInvalidRoutingNumber; err != nil {
		panic(err)
	}
}
