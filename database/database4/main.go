package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	u "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)
func Init()string{
	if err := u.Load();err != nil{
		log.Fatal(err)
	}
	 dbCreds:= fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
		return dbCreds
	
}
func main(){
	var id int
	var name string
	queryTable(id, name)
}
func queryTable(id int, name string){
	dbCreds := Init()
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
