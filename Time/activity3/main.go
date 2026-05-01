package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(whatstheclock())
}
func whatstheclock() string {
	return time.Now().Format(time.ANSIC)
}
