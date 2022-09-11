package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		z -= (math.Pow(z, 2) - x) / (2 * z)
	}
}

func main() {
	fmt.Println(math.Sqrt(2))
}
