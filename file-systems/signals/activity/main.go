package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)
func main(){
	sigs := make(chan os.Signal, 1)
	done := make(chan struct{})
	signal.Notify(sigs, syscall.SIGINT)
	go func(){
		for {
			s := <- sigs
			switch s {
				case syscall.SIGINT:
				fmt.Println()
				fmt.Println("My process has been interrupted. someone might pressed CTRL+C")
				fmt.Println("Some clean up is occuring")
				time.Sleep(5 * time.Second)
				done <- struct{}{}
				return
			}
		}
	}()
	fmt.Println("Program is blocked until a signal is caught")
	<- done
	fmt.Println("Out of here...I'm done")
}