package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func main() {	
	pid, signal := taskTimer()	
	defer psKiller(pid, signal)
	
	
}

func psKiller(pid int ,signal syscall.Signal){// returns a termination signal and pid
    timeOut := 5 * time.Second
    log.Println("Off to sleep it's time to kill this process: ", pid)
   time.Sleep(timeOut)
  	if err := syscall.Kill(pid, signal); err != nil{
		fmt.Println("Error killing process: ", pid, err)
	}
	
}
func taskTimer() (int, syscall.Signal){
	
	signal := syscall.SIGTERM

		cmd := exec.Command("watch","-n", "1", "pstree")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil{
			fmt.Println("Error executing cmd remotely", err)
		}
		pid := cmd.Process.Pid
		log.Println("executing process: ", pid)
	
	return pid, signal

}