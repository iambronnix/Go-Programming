package main
import (
	"fmt"
	"os"
)
func main(){
	args := os.Args
	if len(args) < 2{
		fmt.Println("Usage: go run main.go <ur name>")
		return
	}
	name := args[1]
	greeting:= fmt.Sprintf("Hello %s ... welcome to the cmd", name)
	fmt.Println(greeting)
}