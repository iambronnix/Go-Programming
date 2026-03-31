package main

func (e *errorString) Error() string { // errorString type has a method that implements the error interface
	return e.s // satisfies the requirements, provides Error() method and returns string
}
