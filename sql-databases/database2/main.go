package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	u "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)
func Init2(){
	fmt.Println("~~~~~~~~~~Loading things up~~~~~~~~~")
	time.Sleep(5*time.Second)
}
func Init()string{ //does that env loading thing
	err := u.Load()
	if err != nil{
		log.Fatal(err)
	}
	dbCreds := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	return dbCreds
}
func main(){
	Init2()
	createTable(Init())
	insertTable(Init())
	
}
func createTable(dbCreds string){
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
func insertTable(dbCreds string){
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