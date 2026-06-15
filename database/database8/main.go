package main

import (
	d "github.com/iambronnix/db"
	"fmt"

)

type Users struct{
	id string
	Name string
	Email string
	
}
func main(){
	user1 := Users{id:"10081",Name: "erick", Email: "erickndeto@gmai.com"}
	user2 := Users{id :"10080", Name: "Jayden", Email: "jaydenndeto@gmail.com"}
	myUsers(user1,user2)
	
	fmt.Println("Creating table in a few")
	fmt.Println(createTable())
	
}
func createTable()string{
	db,_ := d.Config()
	createStatement := `
	CREATE TABLE Users(
	Id int NOT NULL,
	Name text NOT NULL,
	Email text NOT NUll,
	)
	WITH (
	OIDS=FALSE
	)
	`
	_, queryErr := db.Query(createStatement)
	if queryErr!=nil{
		panic(queryErr)
	}
	return fmt.Sprintln("Table created successfully")
}
func insertTable(user ...myUsers){
 db, _ := d.Config()
insert, prepareErr := db.Prepare("INSERT INTO Users VALUES(where $1, $2, $3)")	
if prepareErr!= nil{
	panic(prepareErr)
}
for i:=0 ; len(Users);i++{
	_, insertErr := insert.Exec([i]Users...)
}

}

