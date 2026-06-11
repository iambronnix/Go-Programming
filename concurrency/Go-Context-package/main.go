package main

import (
	"context"
	"log"
	"time"
)
func main(){
	r , ctx := make(chan int), context.TODO()//create an integer channel to pass to the counter and a context
	cl,stop := context.WithCancel(ctx)//enables context cancelling
	go countNumbers(cl, r)
	go func(){
		time.Sleep(time.Microsecond*100*3)
		stop()
	}()
	v := <-r
	log.Println(v)
}
func countNumbers(ctx context.Context, r chan int){
	v := 0
	for {
		select {
			case <-ctx.Done()://
			r <- v
			return
			default://if context isn't done we keep counting
			time.Sleep(time.Microsecond*100)
			v++
		}
	}
}