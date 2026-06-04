package main

import (

	"fmt"
   db "github.com/iambronnix/db"
	
)
func main(){
	fmt.Println(truncateTable())
}

func truncateTable()string{
	db, _:= db.Config()
	defer db.Close()
	emptyTable, emptyTableErr := db.Exec("TRUNCATE TABLE test")
	if emptyTableErr !=nil{
		panic(emptyTableErr)
	}
	return fmt.Sprintf("%s",emptyTable)
}