package main

import (
	
	"fmt"
	"time"
	db "github.com/iambronnix/db"

)

func main(){
	createTable()
	time.Sleep(5*time.Second)
	insertTable()
}
func createTable(){
    db , _ := db.Config()
	dbTable := `
	CREATE TABLE Numbers(
	Number int NOT NULL,
	Property text COLLATE pg_catalog."default" NOT NULL
	)
	WITH(
	OIDS = FALSE
	)
	TABLESPACE pg_default;
	ALTER TABLE Number
	OWNER to postgres;
	`
	defer func(){//this recover and ensure insertTable() is executed in case of a panic
		if err := recover();err !=nil{
			fmt.Println("error creating table")
		}
	}()
	defer db.Close()
	_, sqlErr := db.Exec(dbTable)
	if sqlErr != nil{
		panic(sqlErr)//will panic coz proably the table exists
	}else{
		fmt.Println("The table called Number was successfully created")
	}
	
}
func insertTable(){
	db, _ := db.Config()
	defer db.Close()
	insert , insertErr := db.Prepare("INSERT INTO Number VALUES($1, $2)")
	if insertErr != nil{
		panic(insertErr)
	}
	var prop string
	for i := 0; i < 100; i++{
		if i %2 == 0{
			prop = "Even"
		}else{
			prop = "Odd"
		}
		_, err := insert.Exec(i, prop)
		if err == nil{
		fmt.Println("The number:",i,"is:",prop)
		}
		
	}

}
