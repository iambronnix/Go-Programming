package main

import (
	"fmt"
	"sync"
)
var (
	input = make(chan int)//unbuffered channels creates synchronization
	output = make(chan int)
)
func squareWorker(wg *sync.WaitGroup){//sub-squareWorker
	defer wg.Done()//decrements the subworker from waitgroup
	for n := range input{
		output <- n *n//sends squared values through the channel
	}
}
func main(){
	nums := []int{2,4,5,6,7,8,9}
	var wg sync.WaitGroup
	for i := 0; i < 5; i ++{ //creates 5 sub-workers 
		wg.Add(1)
		go squareWorker(&wg)
	}
	go func(){
		for _, n := range nums{//sends all slice values through the channel
			input <- n
		}
		close(input)//closes the channel
	}()
	go func(){
		wg.Wait()
		close(output)
	}()
	for result := range output{
		fmt.Println(result)		
	}
}