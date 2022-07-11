package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64 // reduce error in calculations
	Type() string
}

type Rectangle struct {
	Length, Width float64 // reduce error in calculations
}

type Circle struct {
	Radius float64 // reduce error in calculations
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Type() string {
	return "Rectangle"
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Type() string {
	return "Circle"
}

func main() {
	var s Shape = Rectangle{5.0, 6.0}
	fmt.Println(s.Type(), s.Area())
	var s1 Shape = Circle{3.0}
	fmt.Println(s1.Type(), s1.Area())
}
