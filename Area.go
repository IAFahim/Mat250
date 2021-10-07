package main

import "math"

func WritingNotion() string {
	return "z = f(x,y)"
}

//	Area

func CylinderArea(r, h float64) float64 {
	return math.Pi * r * r * h
}

func ConeArea(r, h float64) float64 {
	return math.Pi * r * ((r * h) / 3)
}

func RectangleArea(x, y float64) float64 {
	return x * y
}

func main() {

}
