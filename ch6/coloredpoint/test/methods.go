package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {

	p := Point{1, 2}
	q := Point{4, 6}
	var origin Point

	// значення-метод
	distanceFromP := p.Distance
	fmt.Println(distanceFromP(q))
	fmt.Println(distanceFromP(origin))

	//вираз-метод
	distance := Point.Distance
	fmt.Println(distance(p, q))
}
func (p Point) Distance(q Point) float64 {
	dX := q.X - p.X
	dY := q.Y - p.Y
	return math.Sqrt(dX*dX + dY*dY)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
