package main

import "fmt"

type Shaper interface {
	Area() string
}
type triangle struct {
	base   float64
	height float64
}
type square struct {
	length float64
	width  float64
}
type circle struct {
	radius float64
}

func main() {
	t := triangle{base: 16, height: 8}
	s := square{length: 4, width: 4}
	c := circle{radius: 14}
	shapeInfo(t, s, c)
}

func shapeInfo(shape ...Shaper) { //shapeInfo function takes a variable number of Shaper interfaces as an argument and calls the Area method of each Shaper interface and prints the result
	for _, s := range shape {
		fmt.Println(s.Area())
	}
}

// Area method for the shape  struct calculates the area  and returns a string representation of the area
func (t triangle) Area() string {
	area := t.base * t.height * 0.5
	return fmt.Sprintf("The area of Triangle is %v", area)
}
func (s square) Area() string {
	area := s.length * s.width
	return fmt.Sprintf("The area of the square is %v", area)
}
func (c circle) Area() string {
	area := c.radius * c.radius * 22 / 7
	return fmt.Sprintf("The Area of the circle is %v", area)
}
