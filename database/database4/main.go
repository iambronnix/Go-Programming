package main

import (
	"fmt"

	db "github.com/iambronnix/db"
)

func main(){
	var id int
	var name string
	queryTable(id, name)
}

func queryTable(id int, name string){
	db,_ := db.Config()
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