package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// func (p *Point) ScaleBy(factor float64) {
// 	p.X *= factor
// 	p.Y *= factor
// }
func (p Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
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

type V int

func (v V) f() V {
	v = 2
	return v
}
func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	distanceFromP := p.Distance
	fmt.Printf("%T\n", distanceFromP)
	fmt.Println(distanceFromP(q))
	var origin Point
	fmt.Println(distanceFromP(origin))
	scaleP := p.ScaleBy
	scaleP(2)
	fmt.Println(p)
}
