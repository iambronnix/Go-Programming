package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"
)
var	processId = make(chan int)


func main() {	
	//synchronous goroutines
	go psKiller()
	executeTask()
	  	
	
}

func psKiller(){// returns a termination signal and pid
	pid := <- processId //receives pid of the current running process inthe terminal
	
	log.Println("Off to sleep it's time to kill this process: ", pid)
   	if err := syscall.Kill(pid,syscall.SIGTERM); err != nil{//sends a termination signal
   fmt.Fprintf(os.Stderr,"%v",err)
	}
	
}
func executeTask(){
	
        cmd := exec.Command(os.Args[0],os.Args[1],os.Args[2])
		cmd.Stdout = os.Stdout
		if err := cmd.Start();err != nil{//starts the process 
			fmt.Fprintf(os.Stderr,"%v",err)		
		}
		pid := cmd.Process.Pid//get process-id of the process being executed inthe terminal	
	    log.Println("executing process: \n", pid)

		time.Sleep(5*time.Second)//psKiller is delayed before receiving pid
			
		processId <- pid//sends pid of running process
		
}

//input := os.Stdin
	//	scanner := bufio.NewScanner(input)
		//for scanner.Scan(){
//scanErr := scanner.Err()
	//	if scanErr == io.EOF{
		//	break
		//}else{
			//continue
		//}
		
//	}