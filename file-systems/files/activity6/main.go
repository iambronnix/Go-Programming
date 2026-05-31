package main

import (
	"errors"
	"io"
	"log"
	"os"
	"strconv"
)
var(
	ErrWorkingFileNotFound = errors.New("Working file doesn't exist")
)
func main(){
	err := createBackup(createFile())
	log.SetFlags(log.Llongfile)
	if err != nil{
		log.Fatal(err)
	}
	workFile,_ := createFile()
	for i := 1; i <=10; i++{
		note  := " " + strconv.Itoa(i)
		err := addNotes(workFile, note)
		if err != nil{
			log.Fatal(err)
		}
	}

}
func addNotes(workFile, notes string)error{
	notes += "\n"
	f, err := os.OpenFile(
		workFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil{
		return err
	}
	defer f.Close()
	if _, err := f.Write([]byte(notes)); err != nil{
		log.Panic(err)
	}
	return nil
}
func createBackup(filesCreated ...string)error{
	workFile := filesCreated[0]
	backupFile := filesCreated[1]
	_, err := os.Stat(workFile)
	   if err != nil{
		if os.IsNotExist(err){
			return err
		}
	}
	working, err := os.Open(workFile)
	if err != nil{
		return err
	}
	content, err := io.ReadAll(working)
	if err != nil{
		return err
	}
	err  = os.WriteFile(backupFile, content, 0644)
	if err != nil{
		log.Fatal(err)
	}
	return nil
}
func createFile()(string, string){
	message := "Notes\n1.Get better at coding\n"
	 files := [2]string{os.Args[1], os.Args[2]}    
       
	for _, i := range files{
		os.Create(i)
		err := os.WriteFile(files[0],[]byte(message),0644)
		if err!=nil{
			log.Fatal(err)
		}

	}
	return files[0], files[1]
}
