package main

import (
	"database/sql"
	d "github.com/iambronnix/db"
	"fmt"
	"log"
	"time"
	_ "github.com/lib/pq"
)
func Init2(){
	fmt.Println("~~~~~~~~~~Loading things up~~~~~~~~~")
	time.Sleep(5*time.Second)
}

func main(){
	Init2()
	createTable()
	insertTable()
	
}
func createTable(){
	dbCreds, err := d.Config()
	if err != nil{
		log.Fatal(err)
	}
	db, err := sql.Open("postgres",dbCreds)
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Println("The connection to the DB was successfully initialised!")
	}
	defer time.Sleep(5 * time.Second)
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
	_, err = db.Exec(dbTable)
	if err != nil{
		panic(err)
	}else{
		fmt.Println("The table was successfully created")
		}
	
}
func insertTable(){
	dbCreds, err := d.Config()
	if err != nil{
		log.Fatal(err)
	}
	db , err := sql.Open("postgres", dbCreds)
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Println("The connection to the DB was successfully initialised!!!")
	}
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