package main

import (
	"fmt"
)

type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
	return 0
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	distanceFromP := p.Distance
	fmt.Println(distanceFromP(q))
	fmt.Printf("%T\n", distanceFromP)
}
