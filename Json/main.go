package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`        //alternative JSON name
	Color  bool `json:"color,omitempty"` // no output if false
	Actors []string
}

var (
	movies = []Movie{
		{
			Title: "Silicon Valley", Year: 2014, Color: false,
			Actors: []string{"bighead", "jarred", "manish", "gilford"}},
		{
			Title: "Mortal combact", Year: 2026, Color: true,
			Actors: []string{"paul newman"}},
	}
	data          = make(chan []byte)
	marshalDone   = make(chan struct{})
	unmarshalDone = make(chan struct{})
)

func main() {

	go marshal()
	go Unmarshal()
	<-marshalDone
	<-unmarshalDone
}

func marshal() {
	defer func() {
		close(data)
		marshalDone <- struct{}{}
	}()
	mashData, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err) //check for errors
	}
	fmt.Printf("%s\n", mashData) //print json type of the data
	data <- mashData
}
func Unmarshal() {
	defer func() {
		unmarshalDone <- struct{}{}
	}()
	var mashData []struct{ Title string } //can add fields to view
	if err := json.Unmarshal(<-data, &mashData); err != nil {
		log.Fatalf("JSON unmarshalling failed: %s", err)

	}
	fmt.Println("==============decoded data===============")
	fmt.Println(mashData)

}
