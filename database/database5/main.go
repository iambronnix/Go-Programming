package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	_ "github.com/lib/pq"
	d "github.com/iambronnix/db"
)
func main(){
	fmt.Println("Initialising the database creds")//just for fun
	time.Sleep(5 * time.Second)//also this
	updateTable()//updates specific data from the database
	deleteTable()//deletes specific data from the database
}
	
func updateTable(){
	db,err := dbConnection()
	defer db.Close()
	fmt.Printf("........Connection to db was successfully initialised........\n")
	updateStatement := `
	UPDATE test(
	SET name = $1
	WHERE id = $2
	)
	`
	updateResult, updateResultErr := db.Exec(updateStatement)
	if err = db.Ping(); err != nil{
		fmt.Println("......Connection lost.....")
	}else{
		fmt.Println("database still connected...good to go")
	}
	if updateResultErr!= nil{
		panic(updateResultErr)
	}
	updatedRecords, updatedRecordsErr := updateResult.RowsAffected()
	if updatedRecordsErr != nil {
		panic(updatedRecordsErr)
	}
	fmt.Println("Number of records updated: ", updatedRecords)
	
}
func deleteTable(){
	db, err := dbConnection()
	if err != nil{
		panic(err)
	}
	defer db.Close()
	fmt.Println("........Connection to db was successfully initialiased.......")
	deleteStatement := `
	DELETE FROM test
	WHERE id = $1
	`
	deleteResult, deleteResultErr := db.Exec(deleteStatement,2)
	if deleteResultErr != nil{
		panic(deleteResultErr)
	}
	deletedRecords, deletedRecordsErr := deleteResult.RowsAffected()
	if deletedRecordsErr != nil{
		panic(deletedRecordsErr)
	}
	fmt.Println("Number of records deleted: ",deletedRecords)
}
func dbConnection()(*sql.DB ,error){
	dbCreds, err := d.Config()
	if err != nil{
		log.Fatal(err)
	}
	db,err := sql.Open("postgres", dbCreds)
	return db, nil
}