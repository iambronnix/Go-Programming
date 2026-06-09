package main

import (
	"fmt"
	"log"
)
func main(){
	ch := make(chan string)
	go greet(ch)
	ch <- "hello erick"
	log.Println(<-ch)
	log.Println(<-ch)
}
func greet(ch chan string){
	msg := <- ch
	ch <- fmt.Sprintf("Thanks for %s",msg)
	ch<- "hello jayden"
}