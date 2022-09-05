package main

import (
	"fmt"
	"math"
)

type point struct {
	x, y float64
}

func newPoint(x, y float64) *point {
	return &point{
		x: x,
		y: y,
	}
}

func distance(a, b *point) float64 {
	return math.Sqrt(math.Pow((b.x-a.x), 2) + math.Pow((b.y-a.y), 2))
}

func main() {
	x := newPoint(5, 3)
	y := newPoint(6, 2)

	fmt.Println(distance(x, y))
}
