package main

import "fmt"

type Point struct {
	X, Y float64
}

func main() {

	r := Point{1, 2}
	r.ScaleBy(3)

	fmt.Println(r)
}

func (p *Point) ScaleBy(factor float64) {

	//fmt.Printf("Type - %T , %v", p, p)

	p.X = p.X * factor
	//p.X *= factor
	p.Y *= factor
}
