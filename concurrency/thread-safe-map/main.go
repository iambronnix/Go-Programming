package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
)
func main(){
	var countMap sync.Map
	numGoroutine := 5
	var wg sync.WaitGroup
	generateAndCount := func(){
		defer wg.Done()
		//generate random numbers per goroutine
		for i := 0 ; i <=1000; i++{
			//generate random numbers btwn 0 and 9
			randomNumber, err := generateRandom(10)
			if err != nil{
				fmt.Println("Error generating random number:",err)
				return
			}
			updateCount(&countMap, randomNumber)
		}
	}
	for i := 0; i <= numGoroutine; i++{
		wg.Add(1)
		go generateAndCount()
	}
	wg.Wait()
	printCounts(&countMap)
	
}
func generateRandom(max int)(int, error){//generates random number in the range of {0, max}
	n, error := rand.Int(rand.Reader,big.NewInt(int64(max)))
	if error != nil{
		return 0, error
	}
	return int(n.Int64()), nil
}
func updateCount(countMap *sync.Map, key int){//updates the count using load and store methods to safely access and update the count map
	count, _ := countMap.LoadOrStore(key, 0)
	countMap.Store(key,count.(int)+1)
}
func printCounts(countMap *sync.Map){//prints the count from the sync.Map contents
	countMap.Range(func(key, value any)bool{
		fmt.Printf("Number %d: Count %d\n", key, value)
		return true
	})
}