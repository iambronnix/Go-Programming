package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	d "github.com/iambronnix/db"
	_ "github.com/lib/pq"
)

func main(){
	createTable()
	time.Sleep(5*time.Second)
	insertTable()
}
func createTable(){
	fmt.Println("Creating db......")
	dbCreds,err := d.Config()
	if err != nil{
		log.Fatal(err)
	}
	db, err := sql.Open("postgres",dbCreds)
	if err != nil{
		log.Fatal(err)
	}
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
	defer func(){
		if err := recover();err !=nil{
			fmt.Println("error creating table")
		}
	}()
	defer db.Close()
	_, err = db.Exec(dbTable)
	if err != nil{
		panic(err)
	}else{
		fmt.Println("The table called Number was successfully created")
	}
	
}
func insertTable(){
	 dbCreds, err := d.Config() 
		if err != nil {
		log.Fatal(err)
	}
	fmt.Println("landed on insertTable() func")//used as a marker for debugging only
	db, err := sql.Open("postgres", dbCreds)
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Ping();err != nil{
		fmt.Println("Connection to the DB lost")
		log.Fatal(err)
	}
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
		_, err = insert.Exec(i, prop)
		if err == nil{
		fmt.Println("The number:",i,"is:",prop)
		}
		
	}

}
