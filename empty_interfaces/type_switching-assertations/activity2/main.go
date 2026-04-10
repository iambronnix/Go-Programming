package main

import "fmt"

func main() {
	var str interface{} = 21
	if v, isValid := str.(string); isValid == true { //type assertion returns two values, the underlying value and a Boolean value
		fmt.Println(v)

	} else { //When the assertion fails, it will return false
		fmt.Println(isValid) //The return value will be the zero value that you are trying to assert to. It also will not panic

	}
}
