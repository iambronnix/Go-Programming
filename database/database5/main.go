package main

import (
	"database/sql"
	"fmt"
	"log"

	d "github.com/iambronnix/db"
)
	
func updateTable(){
	dbCreds, err := d.Config()
	if err != nil{
		log.Fatal(err)
	}
	db, err := sql.Open("postgres", dbCreds)
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Printf("Connection to db was successfully initialised!!!")
	
}