package main

import "log"
func main(){
	ch := make(chan int, 1)//creates a channel
	ch <- 1
	i := <-ch
	log.Println(i)
}