package main

import (
	"fmt"
)

func main() {
	fmt.Println(Sqrt(2))
}

func Sqrt(x float64) float64 {
	z := 1.0
	z -= (z*z - x) / (2 * x)
	return z
}
