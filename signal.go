package main
import "fmt"
func main(){
	greeting("erick hello", 21)
}
func greeting(name string, age int){
	fmt.Printf("%s, ur age is %d\n",name, age)
}