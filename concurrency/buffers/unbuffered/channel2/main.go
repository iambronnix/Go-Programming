package main

import (
	"fmt"
	"sync"
)
func main(){
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ch := make(chan int)
	go readThem(ch, wg)
	ch <-1
	ch <-2
	ch <-3
	ch <-4
	ch <-5
	close(ch)
	wg.Wait()
}
func readThem(ch chan int, wg *sync.WaitGroup){
	for i := range ch {
		fmt.Println(i)
	}
	wg.Done()
}