package main

import (
	"fmt"
	"math"
)
type Point struct{
	x,y float64	
}
func main(){
	p := Point{1,2}
	q := Point{4,6}
	fmt.Println(Distance(p,q))//function call
	fmt.Println(p.Distance(q))//method call
}
func Distance(p, q Point)float64{
	return math.Hypot(q.x - p.x, q.y- p.y)
}
//method of the Point type 
func (p *Point)Distance(q Point)float64{
	return math.Hypot(q.x-p.x,q.y-p.y)
}