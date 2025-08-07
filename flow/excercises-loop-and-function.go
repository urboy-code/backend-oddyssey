package main

import "math"

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}

	if math.Abs(z*z-x) < 1e-10 {
		return z
	}
	return z
}
