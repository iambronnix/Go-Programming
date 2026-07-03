package main

import (
	"encoding/json"
	"fmt"
	"log"
)
type Movie struct{
	Title string
	Year int  `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}
var movies = []Movie{
	{
		Title: "Silicon Valley",Year: 2014, Color: false,
		Actors: []string{"bighead", "jarred", "manish", "gilford"}},
	{
		Title: "Mortal combact",Year: 2026,Color: true,
		Actors: []string{"paul newman"}},
}

func main(){
	data, err := json.MarshalIndent(movies,",", " ")
	if err != nil{
		log.Fatalf("JSON marshaling failed: %s", err)//check for errors 
	}
	fmt.Printf("%s\n",data)//print json type of the data 
}