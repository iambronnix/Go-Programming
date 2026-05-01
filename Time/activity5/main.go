package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	nowToo := now
	time.Sleep(2 * time.Second)
	later := time.Now()
	if now.Equal(nowToo) {
		fmt.Println("The two time variables are equal!!!")
	} else {
		fmt.Println("The two time variables are different!!!")
	}
	if now.Equal(later) {
		fmt.Println("The two time variables are equal!!!")
	} else {
		fmt.Println("The two time variables are different!!!")
	}
}
