package solutions

import (
	"fmt"
	"math"
)

// Task24 - 24.	Разработать программу нахождения расстояния между двумя точками, которые представлены
// в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
func Task24() {
	point1 := NewPoint(0, 0)
	point2 := NewPoint(0, 10)
	fmt.Printf("point 1: %f\npoint 2: %f\ndistance: %f\n", point1, point2, Distance(point1, point2))
}

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func Distance(p1, p2 *Point) float64 {
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}
