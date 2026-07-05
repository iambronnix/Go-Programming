package main

import (
	"fmt"
	"time"
)
func main(){
	//go lines(100*time.Millisecond)
	news(100* time.Millisecond)
}
func lines(sleep time.Duration){
	for {
	for _, i := range `-\|/`{
		fmt.Printf("%c",i)
		time.Sleep(sleep)
	}
	}
}
func news(delay time.Duration){
	for{
		for _, i := range `iamerick`{
			fmt.Printf("%c",i)
		time.Sleep(delay)	
		}
		fmt.Println()
		
		for _, i := range `how are you doing today!!!`{
			fmt.Printf("%c",i)
			time.Sleep(delay)
		}
		fmt.Println()
		
		}
}