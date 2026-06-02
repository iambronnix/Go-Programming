package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/big"

	d "github.com/iambronnix/db"
)
var(
	prop string
	number int64
	primeSum int64
	newNumber int64
)
func main(){
	if _, err := dbConnection();err != nil{
		log.Fatal(err)
	}
	fmt.Println("....Connection to the database established successfully....")
	getPrimeNumbers()
	
}
func getPrimeNumbers(){
	db,_ := dbConnection()
	defer db.Close()
	allTheNumbers := `
	SELECT * FROM Numbers
	`
	numbers, prepareErr := db.Prepare(allTheNumbers)
if prepareErr!= nil{
	panic(prepareErr)
}	
primeSum = 0
result, queryErr := numbers.Query()
if queryErr!= nil{
	panic(queryErr)
}
fmt.Println("The list of prime numbers: \t")
for result.Next(){
	resultErr := result.Scan(&number, &prop)
	if resultErr != nil{
		panic(resultErr)
	}
	if big.NewInt(number).ProbablyPrime(0){
		primeSum += number
		fmt.Print(" ",number)
	}
}
closeErr := numbers.Close()
if closeErr != nil {
	panic(closeErr)
}
}
func dbConnection()(*sql.DB, error){
	dbCreds , dbErr := d.Config()
	if dbErr != nil{
		panic(dbErr)
	}
	db, sqlErr := sql.Open("postgres", dbCreds)
	if sqlErr != nil{
		panic(sqlErr)		
	}
	defer func(){
		if pingErr := db.Ping();pingErr != nil{
			fmt.Println("....Connection to the database lost.....")
		}else{
			fmt.Printf("Connection is stable....check database credentials")
		}
		
	}()
	return db, nil
}