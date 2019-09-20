package main

import (
	"fmt"
	"math"
)

type Circle struct {
	//x, y, r float64
	x float64
	y float64
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := Circle{
		x: 0,
		y: 0,
		r: 5,
	}
	c.x = 10
	c.y = 5

	fmt.Println(c.area())
}
