package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)
func main(){
	fileName, err := os.Create(os.Args[1])
	if err != nil{
		log.Fatal(err)
	}
	defer fileName.Close()
	if _, err = fileName.WriteString("in this activity i'm trying to know if the file exists in this directory"); err != nil{
		log.Println(err)
	}
	var file string
	flag.StringVar(&file, "name", " ", "File name")
	flag.Parse()
	 exist, err := os.Stat(file);
	if err != nil{
		if os.IsNotExist(err){
			log.Printf("%s: File doesn't exist!!!\n",err)
			log.Println(exist)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("file name: %s\nIsDir: %t\nModTime: %v\nMode: %v\nSize:%d\n", exist.Name(), exist.IsDir(), exist.ModTime(), exist.Mode(), exist.Size())
	
	
}