package shape

import (
	"fmt"
)

type Shaper interface {
	area() string
}
type Triangle struct {
	Base   float64
	Height float64
}
type Square struct {
	Length float64
	Width  float64
}
type Circle struct {
	Radius float64
}

func PrintshapeInfo(shape ...Shaper) { //shapeInfo function takes a variable number of Shaper interfaces as an argument and calls the Area method of each Shaper interface and prints the result
	for _, s := range shape {
		fmt.Println(s.area())
	}
}

// Area method for the shape  struct calculates the area  and returns a string representation of the area
func (t Triangle) area() string {
	area := t.Base * t.Height * 0.5
	return fmt.Sprintf("The area of Triangle is %v", area)
}
func (s Square) area() string {
	area := s.Length * s.Width
	return fmt.Sprintf("The area of the square is %v", area)
}
func (c Circle) area() string {
	area := c.Radius * c.Radius * 22 / 7
	return fmt.Sprintf("The Area of the circle is %v", area)
}
