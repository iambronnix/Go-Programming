package main
import (
    "fmt"
    "os"
)
var users = map[string]string{
    "305":"Sue",
    "204":"Bob",
    "631":"Tracy",
    "073":"Jake",
}

func deleteUserdata(id string){
    delete(users, id)
}
func main(){
    if len(os.Args) < 2 {
        fmt.Println("User ID not passed")
        os.Exit(1)
    }
    userID := os.Args[1]
    deleteUserdata(userID)
    fmt.Println("Users:", users)
}