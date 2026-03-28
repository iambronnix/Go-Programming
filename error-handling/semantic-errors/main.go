package main

import "fmt"

func main() { //If our destination is greater than or equal to 2 km, we are going to take a car. If it is less than 2 km,then we will walk to our destination

	km := 2
	if km > 2 {
		fmt.Println("Take the car") //logically incorrect according to the purpose of the program
	} else {
		fmt.Println("Going to walk")
	}
}
