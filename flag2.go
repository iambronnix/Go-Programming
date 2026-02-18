package main
import (
	"fmt"
	"flag"
)
func main(){
	var v bool
	flag.BoolVar(&v, "value", false, "Requires a boolean for the flag.")
	flag.Parse()
	fmt.Println(v)
}