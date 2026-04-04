package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidLastName      = errors.New("invalid last name")
	ErrInvalidRoutingNumber = errors.New("invalid routing number")
)

func main() {
	err := validateInfo("Erick Ndeto", 12345)
	fmt.Println(err)
	err = validateInfo("Jayden Ndeto", 12395)
	fmt.Println(err)

}

func validateInfo(lastName string, routingNumber int) error {
	if lastName != "Erick Ndeto" {
		return ErrInvalidLastName

	}
	if routingNumber != 12345 {
		return ErrInvalidRoutingNumber
	}
	return nil
}
