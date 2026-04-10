package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// json tag is used to specify the name of the field in the JSON data that corresponds to the struct field. In this case, the Name field in the Person struct will be mapped to the "name" field in the JSON data.
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	s := `{"Name":"joe","Age":18}`
	s2 := `{"Name":"Erick","Age":21}`
	p, err := loadPerson(strings.NewReader(s))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	p2, err := loadPerson2(s2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p2)
}

func loadPerson2(s string) (Person, error) {
	var p Person
	//json.NewDecoder is used to create a new JSON decoder that reads from the provided string.
	err := json.NewDecoder(strings.NewReader(s)).Decode(&p)
	// If there is an error during decoding, it returns the error; otherwise, it returns the decoded Person struct.
	if err != nil {
		return p, err
	}
	return p, nil
}
func loadPerson(r io.Reader) (Person, error) {
	var p Person
	err := json.NewDecoder(r).Decode(&p)
	if err != nil {
		return p, err
	}
	return p, nil
}
