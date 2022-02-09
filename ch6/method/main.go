package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

type P int

func (p *P) f() P {
	*p = 2
	return *p
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))
	fmt.Printf("%T\n", Distance)
	fmt.Printf("%T\n", p.Distance)
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
	var a P = 1
	fmt.Println(a.f())
	fmt.Println(a)
	// 	distanceFromP := p.Distance
	// 	fmt.Println(distanceFromP(q))
	// 	fmt.Printf("%T\n", distanceFromP)
}
