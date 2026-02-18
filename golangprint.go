package main
import "fmt"
func main(){
    greetings("hello erick", 20)
}

func greetings(name string, age int){
    fmt.Printf("%s i'm %d years old. \n", name, age)
}