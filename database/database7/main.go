package main

import (
	"database/sql"
	"fmt"

	d "github.com/iambronnix/db"
)
func main(){
	fmt.Println(truncateTable())
}
func dbConnection()(*sql.DB,error){
	dbCreds, dbErr := d.Config()
	if dbErr != nil{
		panic(dbErr)
	}
	db, sqlErr := sql.Open("postgres", dbCreds)
	if sqlErr != nil{
		panic(sqlErr)
	}
	defer func(){
		if pingErr := db.Ping();pingErr != nil{
			fmt.Println("......Connection to the database lost....")
		}else{
			fmt.Println("Connection is stable ..... check database credentials")
		}
	}()
	return db, nil
}
func truncateTable()string{
	db, _:= dbConnection()
	defer db.Close()
	emptyTable, emptyTableErr := db.Exec("TRUNCATE TABLE test")
	if emptyTableErr !=nil{
		panic(emptyTableErr)
	}
	return fmt.Sprintf("%s",emptyTable)
}