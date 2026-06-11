package main
import "fmt"
func main(){
	ch := make(chan int)//unbuffered channel
	go readThem(ch)
	ch <-1
	ch <-2
	ch <-3
}
func readThem(ch chan int){
	for{
		fmt.Println(<-ch)
	} 
}