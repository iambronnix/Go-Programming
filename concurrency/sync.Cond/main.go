package main

import (
	"fmt"
	"sync"
	"time"
)
type Workqueue struct{
	cond *sync.Cond
	maxSize int
	workItems []string
}
func main(){
	var wg sync.WaitGroup
	workQueue := newWorkQueue(3)
	wg.Add(1)
	go func(){
		defer wg.Done()
		for i := 1; i <=5;i++{
			workItem := fmt.Sprintf("WorkItem %d", i)
			workQueue.enqueue(workItem)
			fmt.Printf("Enqueued: %s\n",workItem)
			time.Sleep(time.Second)
		}
	}()
	wg.Add(1)
	go func(){
		defer wg.Done()
		for i := 1; i <=5 ; i ++{
			workItem := workQueue.dequeue()
			fmt.Printf("Dequeued: %s\n", workItem)
			time.Sleep(2*time.Second)
		}
	}()
	wg.Wait()
}
func newWorkQueue(maxSize int)*Workqueue{
	return &Workqueue{
		cond: sync.NewCond(&sync.Mutex{}),
		maxSize: maxSize,
		workItems: make([]string, 0),
	}
}
func (wq *Workqueue)enqueue(item string){
	wq.cond.L.Lock()
	defer wq.cond.L.Unlock()
	for len(wq.workItems) != 0{//unlocks L resource  
		wq.cond.Wait()
	}
	wq.workItems = append(wq.workItems,item)
	wq.cond.Signal()
}
func (wq *Workqueue) dequeue()string{
	wq.cond.L.Lock()
	defer wq.cond.L.Unlock()
	for len(wq.workItems) == 0{
		wq.cond.Wait()
	}
	item := wq.workItems[0]
	wq.workItems = wq.workItems[1:]
	wq.cond.Signal()
	return item

}