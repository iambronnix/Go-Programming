type error interface { //error is an interface type
	Error() string
}

//to satisfy the error interface, only :metheod name,Error() and Error() method to return string