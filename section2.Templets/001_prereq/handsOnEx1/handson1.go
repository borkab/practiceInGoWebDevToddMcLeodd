package main

import "math"

type square struct {
	//weight float32
	Height float64
}

type circle struct {
	Radius float64
}

func (s square) Area() float64 {
	return s.Height * s.Height
}
func (c circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	Area() float64
}

func Info(s Shape) float64 {
	return s.Area()
}
