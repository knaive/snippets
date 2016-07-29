package interfaceTest

import (
	"fmt"
	"math"
)

func main() {
	c1 := Circle{1.0}
	c2 := Circle{1.0}
	r := Retangle{2.0, 3.0}

	fmt.Println(totalArea(&c1, &c2, &r))
}

type Shape interface {
	area() float64
}

type Circle struct {
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Retangle struct {
	length float64
	width  float64
}

func (r *Retangle) area() float64 {
	return r.length * r.width
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		t := s.area()
		area += t
	}
	return area
}
