package main

import (
	"fmt"
	"math"
)

// Shape represents a geometrical Shape
type Shape interface {
	area() float64
}

// Point2D is a 2D space point with 2 coordinates
type Point2D struct {
	x, y float64
}

// Distance return the distance between two 2D points
func Distance(a, b Point2D) float64 {
	return math.Sqrt((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y))
}

// Circle is a centre and a radius
type Circle struct {
	centre Point2D
	radius float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s *Square) area() float64 {
	return (Distance(s.x, s.y) * Distance(s.x, s.y)) / 2
}

// Square is defined by a diagonal : two 2D space points
type Square struct {
	x, y Point2D
}

func main() {
	c := Circle{Point2D{0, 0}, 2}
	fmt.Println("c area : ", c.area())
	s := Square{Point2D{0, 1}, Point2D{1, 0}}
	fmt.Println("s area : ", s.area())
}
