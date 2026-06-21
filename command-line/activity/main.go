package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"
)
var processId = make(chan int)

func main() {
//synchronous process	
	go psKiller()
     executeTask()	
	
}

func psKiller(){// sends termination signal to the process running inthe background
	pid := <- processId//receives pid of the current running process
	
  	if err := syscall.Kill(pid, syscall.SIGTERM); err != nil{//sends termination signal to pid
		fmt.Fprintf(os.Stderr,"%v",err)
	}
	
}
func executeTask(){
	
		cmd := exec.Command(os.Args[0],os.Args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil{//starts the process
			fmt.Fprintf(os.Stderr,"%v", err)
			
		}
		pid := cmd.Process.Pid
		
		log.Println("executing process: \n", pid)
		time.Sleep(5*time.Second)
		processId <- pid//sends pid of running process
		
	
	

}
//can try with ./main lsof