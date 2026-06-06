package main

import (
	"log"
	"sync"
)
func main(){
	s1 := 0
	wg := &sync.WaitGroup{}//creates a wait group and passes later to sum()
	wg.Add(1)//adds sum to the wait group
	go sum(1,100,wg,&s1)
	wg.Wait()//waits until go routine executes sum()
	log.Println(s1)
	
}
func sum(from , to int, wg *sync.WaitGroup, res *int ){//wg and res point to waitgroups and int 
	*res = 0
	for i := from; i<=to; i++{
		*res += i
	}
	wg.Done()
}
