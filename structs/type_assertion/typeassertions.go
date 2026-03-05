package main

import (
	"errors"
	"fmt"
)

func doubler(v interface{}) (string, error) {
	if i, ok := v.(int); ok {
		return fmt.Sprint(i * 2), nil
	}
	if s, ok := v.(string); ok {
		return s + s, nil
	}
	return "", errors.New("Unsupported type passed")
}

func main() {
	res, _ := doubler(7)
	fmt.Println("7:", res)
	res, _ = doubler("Go")
	fmt.Println("Go:", res)
	_, err := doubler(true)
	fmt.Println("true:", err)
}
