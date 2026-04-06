package main

import (
	"errors"
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
	d := directDeposit{
		firstName:        "Erick",
		lastName:         "Ndeto",
		bankName:         "NCBA",
		routingNumber:    23,
		accountingNumber: 228993988,
	}
	if err := d.validateLastName(); err != nil {
		d.reportPanic(err)
	}
	if err := d.validateRoutingNumber(); err != nil {
		d.reportPanic(err)
	}

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
func (d *directDeposit) reportPanic(err error) {
	if ErrInvalidLastName != nil || ErrInvalidRoutingNumber != nil {

		panic(err)
	}
}
