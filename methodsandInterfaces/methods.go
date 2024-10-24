package main

import (
	fmt
   "math"
)

type Point struct {
	X, Y float64
}

type Line struct {
	Begin, End Point
}

func (l Line) Distance() float64 {
	return math.Hypot(l.End.X-l.Begin.X, l.End.y-l.Begin.Y)
}

func main() {
	side := Line{Point{1, 2}, Point{4, 6}}

	fmt.Println(side.Distance())
}