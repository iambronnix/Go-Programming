package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
	d "github.com/iambronnix/db"
)
func main(){
	fmt.Println("Initialising the database creds")//just for fun
	if _, dbErr := dbConnection();dbErr!=nil{
		log.Fatal(dbErr)
	}
	updateTable()//updates specific data from the database
	deleteTable()//deletes specific data from the database
}
	
func updateTable(){
	db,_ := dbConnection()
	defer db.Close()
	updateStatement := `
	UPDATE test(
	SET name = $1
	WHERE id = $2
	)
	`
	updateResult, updateResultErr := db.Exec(updateStatement)
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
	db,_:= dbConnection()

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
	db,sqlErr := sql.Open("postgres", dbCreds)
	if sqlErr!= nil{
		log.Fatal(sqlErr)
	}
	return db, nil
}