package main

import (
	"fmt"
	"math"
)

// Point x,y
type Point struct {
	X, Y float64
}

// ColoredPoint is a Point with a color
type ColoredPoint struct {
	Point
	Color string
}

// DistanceFrom calculates distance
func (p *Point) DistanceFrom(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Scale scales the point
func (p *Point) Scale(factor float64) {
	if p == nil {
		return
	}
	p.X = p.X * factor
	p.Y = p.Y * factor
}

func main() {
	p1 := ColoredPoint{Point{1, 3}, "red"}
	scaleFn := p1.Scale
	fmt.Println(p1)
	scaleFn(2)
	fmt.Println(p1)
}
