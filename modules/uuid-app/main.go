package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New() //generate a new UUID using external package
	fmt.Printf("Generated UUID: %s\n", id)
}
