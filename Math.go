package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Math test")
	var x, y, z int
	x = 1
	y = 2
	z = x + y
	fmt.Print(math.Sqrt(float64(z)))
}
