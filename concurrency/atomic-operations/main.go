package main

import (
	"log"
	"sync"
	"sync/atomic"
)
func main(){
	s1 := int32(0)
	wg := sync.WaitGroup{}
	wg.Add(4)//adds all the goroutines into the waitgroup
	go sum(1, 25, &wg, &s1)//go routines each access the common variable s1 in turns
	go sum(26, 50, &wg, &s1)
	go sum(51, 75, &wg, &s1)
	go sum(76, 100, &wg, &s1)
	wg.Wait()//waits until all the goroutines are executed
	log.Println(s1)
}
func sum(from , to int, wg *sync.WaitGroup, res *int32){
	//if you remove sync/atomic import and change to 
	// *res = *res + int32(i)
	// the test fails
	for i := from;i <=to; i++{
		atomic.AddInt32(res, int32(i))//adds value i to res and updates the new value
	}
	wg.Done()//decrements the waitgroup goroutines
}