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
	d := directDeposit{
		firstName:        "erick",
		lastName:         "",
		bankName:         "ncba",
		routingNumber:    323,
		accountingNumber: 9283998399383,
	}
	if err := d.validateRoutingNumber(); err != nil {
		d.reportPanic(err)
	}
	if err := d.validateLastName(); err != nil {
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
	defer func() { //recovering from the panic
		if r := recover(); r != nil {

			fmt.Printf("%v\n%v", ErrInvalidLastName, ErrInvalidRoutingNumber)
			fmt.Printf("\n**************************************************\n")
			fmt.Printf("\nfirst name: %v\nlastname: %v\nbankname: %v\nRoutingnumber: %v\naccountNumber: %v\n", d.firstName, d.lastName, d.bankName, d.routingNumber, d.accountingNumber)
		}
	}()
	if ErrInvalidRoutingNumber != nil || ErrInvalidLastName != nil {
		panic(err)
	}
}
