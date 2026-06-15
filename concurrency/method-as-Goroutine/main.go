package main

import (
	"fmt"
	"sync"
)
type Worker struct{
	in, out chan int
	sbw int//sbw:subworker
	mtx *sync.Mutex
}
func main(){
	mtx := &sync.Mutex{}
	in , out := make(chan int, 100), make(chan int)
	wrNum, wrk := 10, Worker{in: in, out: out,mtx: mtx}
	for i := 0; i <=wrNum;i++{//call readThem() method wrNum times
		wrk.readThem()
	}
	for i := 1; i <= 100; i++{//send numbers to be summed to the channel
		in <- i
	}
	close(in)//close the channel to notify all the numbers have been sent
	res := wrk.gatherResult()
	fmt.Println(res)
}
func (w *Worker)gatherResult()int{
	total := 0
	wg := &sync.WaitGroup{}//create a waitgroup instance add routine to the waitgroup
	wg.Add(1)
	go func(){
		for i := range w.out{
			total += i
		}
		wg.Done()
	}()
	wg.Wait()
	return total
	
}
func (w *Worker) readThem(){
	w.sbw++//increment number of subworker instances
	go func(){
		partial := 0
		for i := range w.in{
			partial += i
		}
		w.out <- partial
	w.mtx.Lock()//locked the routine,reduced the counter on the sub-workers safely
		w.sbw--
		if w.sbw == 0{//incase all workers have terminated, close  the output channel
			close(w.out)
		}
		w.mtx.Unlock()
		
	}()
}