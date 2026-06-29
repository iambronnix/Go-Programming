package main

import (
	"log"
	"net/http"
)
type hello struct{}
func(h hello) ServeHTTP(w http.ResponseWriter, r *http.Request){
	msg := "</>Hello World</h1>"
	w.Write([]byte(msg))
}
func main(){
	http.HandleFunc("/chapter1", func(w http.ResponseWriter, r *http.Request){
		msg := "chapter1"
		w.Write([]byte(msg))
	})
	http.Handle("/",hello{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}