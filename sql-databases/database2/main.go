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
	db, err := sql.Open("postgres", dbCreds)
	if err != nil{
		panic(err)
	}else{
		fmt.Println("The connection to the DB was successfully initialized!!!")
	}
	DBCreate := `
	CREATE TABLE public.test(
	id int,
	name character varying COLLATE pg_catalog."default"
	)
	WITH(
	OIDS = FALSE
	)`
	_, err = db.Exec(DBCreate)
	if err != nil{
		panic(err)
	}else{
		fmt.Println("The table was successfully created!!!")
	}
}