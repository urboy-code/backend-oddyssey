package main

import "math"

func pow(x, n, y float64) float64{
	if x := math.Pow(x, n); x < y {
		return x
	}
	return y
}