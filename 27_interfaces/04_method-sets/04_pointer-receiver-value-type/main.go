package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type shape interface {
	area() float64
}

func (c *circle) area() float64 { // pointer receiver
	return math.Pi * c.radius * c.radius
}

func info(z shape) {
	fmt.Println("Area is", z.area())
}

func main() {
	x := circle{5}
	info(x) // error because a value cannot be passed to a pointer receiver
}
