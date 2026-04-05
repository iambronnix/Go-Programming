package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidLastName      = errors.New("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n\ninvalid last name")
	ErrInvalidRoutingNumber = errors.New("invalid routing number\n")
)

type directDeposit struct {
	firstName     string
	lastName      string
	bankName      string
	routingNumber int
	accountNumber int
}

func main() {
	d := directDeposit{firstName: "Erick", lastName: "Ndeto", bankName: "ncba bank", routingNumber: 726535532, accountNumber: 87263838873}
	k := directDeposit{firstName: "Jayden", lastName: "", bankName: "absa bank", routingNumber: 8788533344, accountNumber: 9828298873}
	j := directDeposit{firstName: "Meozy", lastName: "Mwendwa", bankName: "equity bank", routingNumber: 32, accountNumber: 8728879883}
	p := directDeposit{firstName: "Yaah", lastName: "", bankName: "I&M bank", routingNumber: 87, accountNumber: 87263838873}
	if err := d.validateRoutingNumber(); err != nil {
		d.report(err)
	}
	if err := d.validateLastName(); err != nil {
		d.report(err)
	}
	if err := k.validateRoutingNumber(); err != nil {
		k.report(err)
	}
	if err := k.validateLastName(); err != nil {
		k.report(err)
	}
	if err := j.validateRoutingNumber(); err != nil {
		j.report(err)
	}
	if err := j.validateLastName(); err != nil {
		j.report(err)
	}
	if err := p.validateRoutingNumber(); err != nil {
		p.report(err)
	}
	if err := p.validateLastName(); err != nil {
		p.report(err)
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
func (d *directDeposit) report(error) string {
	if ErrInvalidLastName != nil || ErrInvalidRoutingNumber != nil {
		fmt.Printf("%v\n%v", ErrInvalidLastName, ErrInvalidRoutingNumber)
		fmt.Printf("**************************************************\n")
		fmt.Printf("\nfirst name: %s\nlastname: %s\nbankname: %s\nRoutingnumber: %d\naccountNumber:%d\n", d.firstName, d.lastName, d.bankName, d.routingNumber, d.accountNumber)
	}

	return fmt.Sprintf("\nfirst name: %s\nlastname: %s\nbankname: %s\nRoutingnumber: %d\naccountNumber:%d", d.firstName, d.lastName, d.bankName, d.routingNumber, d.accountNumber)

}
