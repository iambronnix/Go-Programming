package main
import (
	"fmt"
	"database/sql"
)
func main(){
	db, err := sql.Open("postgres", "user=bronnix-nerd password=Start123 host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
	if err != nil{
		panic(err)
	}else{
		fmt.Println("The connection to the DB was successfully initialised!!!")
	}
	connectivity := db.Ping()
	if connectivity != nil{
		panic(err)
	}else{
		fmt.Println("Good to go!")
	}
	db.Close()
	
}