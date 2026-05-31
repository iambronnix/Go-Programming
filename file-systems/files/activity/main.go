package main

import (
	"fmt"
	"os"
	"time"
)
func main(){
	f, err := os.Create(cmd())
	if err != nil{
		panic(err)
	}
	defer f.Close()
	f.Write([]byte("using write() to write some text into the file.\n"))
	f.WriteString("usig writestring() to write into the file created.\n")
}
func cmd()string{
	time.Sleep(1*time.Second)
	filename := os.Args[1]
	fmt.Printf("file %s created\n", os.Args[1])
	return filename
}