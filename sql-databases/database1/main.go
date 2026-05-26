package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	u "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)
func main(){
	err := u.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connStr:= fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)		
	db, err := sql.Open("postgres", connStr)
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Println("The connection to the DB was successfully initialised!!!")
	}
	connectivity := db.Ping()
	if connectivity != nil{
		log.Fatal(err)
	}else{
		fmt.Println("Good to go!")
	}
	db.Close()
	
}