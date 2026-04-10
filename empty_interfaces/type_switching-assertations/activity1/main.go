package main

import (
	"fmt"
	"strings"
)

func main() {
	var str interface{} = "some string"
	v := str.(string) //asserts that str is of the string type and assigns it to the variable v
	fmt.Println(strings.Title(v))

}
