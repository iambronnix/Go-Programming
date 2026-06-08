package main

import "fmt"
func main(){
	ch := make(chan string,1)
	name:= myName(ch)
	fmt.Println(<-name)
	
}

func myName(msg chan string)chan string{
	msg <- "erick ndeto"
	return msg		
}