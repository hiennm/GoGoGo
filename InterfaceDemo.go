package main

import (
	"math"
)

//Abser interface
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := nest
}

type MyFloat float64

//Abs function
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//Vertex struct
type Vertex struct {
	X, Y float64
}

//Abs function: calculate the absolute distance
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
