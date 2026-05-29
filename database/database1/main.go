package main

import (
	"database/sql"
	"fmt"
	"log"

	d "github.com/iambronnix/db"
	_ "github.com/lib/pq"
)

func main() {
	dbCreds, err := d.Config()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("postgres", dbCreds)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("The connection to the DB was successfully initialised!!!")
	}
	connectivity := db.Ping()
	if connectivity != nil {
		log.Fatal("error with connection")
	} else {
		fmt.Println("Good to go!")
	}
	db.Close()

}
