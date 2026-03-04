package main
import (
    "fmt"
    "os"
    "flag"
)
var (users = map[string]string{}
     v string
    id int)
     

func main(){
    if len(os.Args) < 2 {
        fmt.Println("No user passed")
        os.Exit(1)
    }
    for i := 0; i < 5; i++ {
        flag.IntVar(&v, "user", 0, "user not passed")
        flag.Parse()
        id = v
        append(users, id)
    }
}