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
	d := directDeposit{firstName: "Erick", lastName: "Ndeto", bankName: "NCBA", routingNumber: 34443, accountingNumber: 228993988}
	err := d.validateLastName()
	if err != nil {
		fmt.Println(err)
	}
	err = d.validateRoutingNumber()
	if err != nil {
		fmt.Println(err)
	}
	d.report(err)

}

func (d *directDeposit) validateRoutingNumber() error {
	if d.routingNumber < 100 {
		return ErrInvalidRoutingNumber
	}
	return nil
}
func (d *directDeposit) validateLastName() error {
	if len(d.lastName) == 0 {
		return ErrInvalidLastName
	}
	return nil
}
func (d *directDeposit) report(err error) {
	if ErrInvalidLastName != nil || ErrInvalidRoutingNumber != nil {
		panic(err)
	}

}
