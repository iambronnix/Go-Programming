package main 
import (
    "fmt"
    "time"
    "flag"
)
var(
    dayborn time.Weekday
)
func main(){
    flag.StringVar(&dayborn, "hbd",Monday, "Dayborn is required with a -hbd flag")
    flag.Parse()
}
func hbdmessage(string){
    switch dayborn{
    case time.Monday:
        fmt.Println("Monday's child is fair of face")
    case time.Tuesday:
        fmt.Println("Tuesday's child is full of grace")
    case time.Wednesday:
        fmt.Println("Wednesday child is full of woe")
    case time.Thursday:
        fmt.Println("Thursday child has far to go")
    case time.Friday:
        fmt.Println("Friday's child is loving and giving")
    case time.Saturday:
        fmt.Println("Saturday's child works hard for a living")
    case time.Sunday:
        fmt.Println("Sunday's child is bonny and blithe")
    default:
        fmt.Println("Error, it seems you weren't born, hahaha")
    }
}