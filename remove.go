package main
import (
    "fmt"
    "os"
)

var action = map[string]string{
    "1":"Good",
    "2":"Good",
    "3":"Bad",
    "4":"Good",
    "5":"Good",
} 

func main(){
    if len(os.Args) < 2 {
        fmt.Println("No ID passed\n",action)
        os.Exit(1)
    }
   var id string = os.Args[1]
   delete(action, id) 
   fmt.Println("Updated list",action)
}