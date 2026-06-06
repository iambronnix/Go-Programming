package main

import "fmt"

func main(){
	a := 0
	k := next(&a)//each func updates v independently without cordination
	j := next(&a)
	m := next(&a)
	n := next(&a)
	fmt.Println(k, j,m,n)//the output doesn't guarantee there won't be a race condition on v
	
}
func next(v *int)int{
	c := *v 
	*v = c +1 
	return *v
}