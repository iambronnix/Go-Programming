package main

import (
	"fmt"
	db "github.com/iambronnix/db"
)


func main(){
	createTable()
	insertTable()
	
}
func createTable(){
 db ,_ := db.Config()
	defer func(){//handles panic incase of table exists
		err := recover()
		if err != nil{
			fmt.Println("Error in creating a table")
		}
	}()
	defer db.Close()//closes the db just before exit
	dbTable := `
	CREATE TABLE publictest(
	id int,
	name character varying COLLATE pg_catalog."default"
	)
	WITH(
	OIDS = FALSE
	)
	`
	_, err := db.Exec(dbTable)
	if err != nil{
		panic(err)
	}else{
		fmt.Println("The table was successfully created")
		}
	
}
func insertTable(){
    db, _ := db.Config()
	defer db.Close()
	insert, err := db.Prepare("INSERT INTO publictest VALUES ($1, $2)")
	if err != nil{
		panic(err)
	}
	_, err = insert.Exec(2,"second")//continue inserting...
	if err != nil{
		panic(err)
	}
	fmt.Println("The value was successfully inserted!")
}