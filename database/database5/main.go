package main

import (
	"fmt"
	db "github.com/iambronnix/db"
)
func main(){
   	updateTable()//updates specific data from the database
	deleteTable()//deletes specific data from the database
}
	
func updateTable(){
	db,_ := db.Config()
	defer db.Close()
	updateStatement := `
	UPDATE test
	SET name = $1
	WHERE id = $2
	`
	updateResult, updateResultErr := db.Exec(updateStatement,"well",2)
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
	db,_:= db.Config()

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