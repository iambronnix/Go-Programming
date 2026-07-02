package select_test

import (
	"fmt"
	"time"
)
var (
	chan1 = make(chan int)
	chan2 = make(chan int)
)
func processMessages(){
	delay := time.After(1*time.Second)
 for {
   select {//listens from any of the channels 1 second
   case msg1 := <-chan1:
   fmt.Println("Received from chan1", msg1)
   case msg2 := <-chan2:
   fmt.Println("Received from chan2", msg2)
   case <- delay:
   fmt.Println("No messages received in :",delay)
   }
 }
 }