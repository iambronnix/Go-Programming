package main

import (
	"log"
	"sync"
)
func main(){
	res := sum(100,1,100)
	log.Println(res)
}
func worker(in chan int, out chan int){//worker pod 
	sum := 0
	for i := range in{//a loop that ranges over the channel
		sum += i
	}
	out <- sum
	<- in
}
func sum(workers,from, to int)int{
	wg := &sync.WaitGroup{}
	out:=make(chan int, workers)
	in:=make(chan int,5)
	go func(){
		for i := 0; i <= workers; i++{//runs requested number of workers
		wg.Add(i)
		go worker(in,out)
	}
	wg.Done()
	for i := from; i <=to; i++{
		in <- i//sends all numbers to be summed to the in channel,which will distribute the numbers across all Goroutines.
	}
	close(in)//notifies worker func numbers to sum are finished
	}()
	wg.Wait()
	sum := 0
	for i := 0 ; i <= workers; i ++{//performs sum of the partials
		sum += <- out
	}
	close(out)
	return sum
}