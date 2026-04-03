package main

import (
	"errors"
	"fmt"
)

func main() {
	a()
	fmt.Println("This line will now get printed from main() function")
}
func a() {
	b("good-bye") // this will cause a panic in func b() and the defer() in func b() will recover from the panic and print the error message, then the execution will continue in func a() and print "Back in function a()"
	fmt.Println("Back in function a()")
}
func b(msg string) {
	defer func() { // this defer() will recover from the panic in func b() and print the error message, then the execution will continue in func a() and print "Back in function a()"
		if r := recover(); r != nil {
			fmt.Println("error in func b()\n", r)
		}
	}()
	if msg == "good-bye" {
		panic(errors.New("something went wrong\n")) // this will cause a panic in func b() and the defer() in func b() will recover from the panic and print the error message, then the execution will continue in func a() and print "Back in function a()"
	}
	fmt.Print(msg) // this line will not get executed because the panic will occur before it, and the defer() will recover from the panic and print the error message, then the execution will continue in func a() and print "Back in function a()"
}
