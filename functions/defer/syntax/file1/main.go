package main

import "fmt"

func main() {
	defer done()
	fmt.Println("starting the main function")
	fmt.Println("End of main func")
}
func done() {
	fmt.Println("func done() got called finally")
}
