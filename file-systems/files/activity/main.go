package main

import (
	"fmt"
	"os"
	"time"
)
func main(){
	f, err := os.Create(cmd())
	if err != nil{
		fmt.Println("Error creating file", err)
	}
	defer f.Close()
}
func cmd()string{
	time.Sleep(1*time.Second)
	filename := os.Args[1]
	fmt.Printf("file %s created\n", os.Args[1])
	return filename
}