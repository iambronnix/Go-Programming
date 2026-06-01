package main

import (
	"database/sql"
	"fmt"
	"log"
    d "github.com/iambronnix/db"	
	_ "github.com/lib/pq"
)

func main(){
	var id int
	var name string
	queryTable(id, name)
}
func queryTable(id int, name string){
	dbCreds, err := d.Config()
	if err != nil{
		log.Fatal(err)
	}
	db, err := sql.Open("postgres",dbCreds)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connection to postgres was successfully initialized")
	defer db.Close()
	 rows, err := db.Query("SELECT * FROM Number")
		if err != nil{
		panic(err)
	}
	for rows.Next(){
		if err := rows.Scan(&id, &name); err!= nil{
			panic(err)
		}
		fmt.Printf("Retrived data from db: %d %s\n", id, name)
	}
	if err = rows.Err(); err != nil{
		panic(err)
	}
	if err = rows.Close(); err != nil{
		panic(err)
	}
	
}
