package main

import (
	"activity/pkg/shape"
)

func main() {
	t := shape.Triangle{Base: 12, Height: 8}
	s := shape.Square{Length: 5, Width: 8}
	c := shape.Circle{Radius: 7}
	shape.PrintshapeInfo(t, s, c)
}
