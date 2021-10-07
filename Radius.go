package main

import "math"

func SpareRadius(x, y, z float64) float64 {
	return math.Sqrt(x*x + y*y + z*z)
}

func planeRadius(x, y, z float64) float64 {
	return x + y + z
}

func main() {

}
