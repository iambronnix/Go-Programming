package main

import (
	"fmt"
	"math"

)
type Point struct{
	x,y float64	
}
type Path []Point//a Path is journey connecting the points with straight lines
func main(){
	p := Point{1,2}
	q := Point{4,6}
	perim := Path{
		{1,1},
		{5,1},
		{5,4},
		{1,1},
	}
	fmt.Println(Distance(p,q))//function call
	fmt.Println(p.Distance(q))//method call
	fmt.Println(perim.Distance())
}
func Distance(p, q Point)float64{
	return math.Hypot(q.x - p.x, q.y- p.y)
}
//method of the Point type 
func (p *Point)Distance(q Point)float64{
	return math.Hypot(q.x-p.x,q.y-p.y)
}
func(path Path) Distance() float64{
	sum := 0.0
	for i := range path{
		if i > 0{
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}