package main

import (
	"fmt"
	"math/big"
	db "github.com/iambronnix/db"
)

var (
	prop      string
	number    int64
	primeSum  int64
	newNumber int64
)

func main() {
	getPrimeNumbers()
	removeEvenNumbers()
	updateRecords()

}
func getPrimeNumbers() {
	db, _ := db.Config()
	defer db.Close()
	allTheNumbers := `
	SELECT * FROM number
	`
	numbers, prepareErr := db.Prepare(allTheNumbers)
	if prepareErr != nil {
		panic(prepareErr)
	}
	primeSum = 0
	result, queryErr := numbers.Query()
	if queryErr != nil {
		panic(queryErr)
	}
	fmt.Println("The list of prime numbers: \t")
	for result.Next() {
		resultErr := result.Scan(&number, &prop)
		if resultErr != nil {
			panic(resultErr)
		}
		if big.NewInt(number).ProbablyPrime(0) {
			primeSum += number
			fmt.Print(" ", number)
		}
	}
	closeErr := numbers.Close()
	if closeErr != nil {
		panic(closeErr)
	}
}
func removeEvenNumbers() {
	db, _ := db.Config()
	defer db.Close()
	remove := "DELETE FROM number WHERE Property=$1"
	removeResult, removeErr := db.Exec(remove)
	if removeErr != nil {
		panic(removeErr)
	}
	modifiedRecords, modifiedRecordErr := removeResult.RowsAffected()
	if modifiedRecordErr != nil {
		panic(modifiedRecordErr)
	}
	fmt.Println("The number of rows removed : ", modifiedRecords)
	fmt.Println("Updating numbers....")
}
func updateRecords() {
	db, _ := db.Config()
	defer db.Close()
	update := `UPDATE number SET Number=$1 WHERE Number = $2 AND Property=$3`
	allTheNumbers := `SELECT * FROM Numbers`
	numbers, prepareErr := db.Prepare(allTheNumbers)
	if prepareErr != nil {
		panic(prepareErr)
	}
	results, resultErr := numbers.Query()
	if resultErr != nil {
		panic(resultErr)
	}
	for results.Next() {
		scanErr := results.Scan(&number, &prop)
		if scanErr != nil {
			panic(scanErr)
		}
		newNumber = number + primeSum
		_, execErr := db.Exec(update, newNumber, number, prop)
		if execErr != nil {
			panic(execErr)
		}
	}
	if closeErr := numbers.Close(); closeErr != nil {
		panic(closeErr)
	}
	fmt.Println("The exection is now complete...")

}
