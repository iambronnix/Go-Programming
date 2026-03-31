package main

import (
	"fmt"
)

func main() {
	es := errors.errorString{} //trying to access an unexported errorString type and it's field
	es.s = "slacker"
	fmt.Println(es)
}
