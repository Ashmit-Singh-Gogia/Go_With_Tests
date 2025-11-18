package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}
type Triangle struct {
	Base, Height float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
