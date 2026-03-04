package main
import (
	"fmt"
)

type FileParser interface {
	loadEmployee(s string)
	loadFromFile(f *os.File)
	loadFromHttpRequest(r *Request)

}