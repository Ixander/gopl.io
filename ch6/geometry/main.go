package main

import (
	"GoBook/ch6/geometry/geo"
	"fmt"
)

func main() {

	p := geo.Point{X: 1, Y: 2}
	q := geo.Point{X: 1, Y: 5}

	res := p.Distance(q)
	fmt.Println(res)

	fmt.Println(geo.Distance(p, q))

	perim := geo.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
}
