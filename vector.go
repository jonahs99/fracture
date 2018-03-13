package main

import "math"

// Vector2 is a 2d vector
type Vector2 struct {
	x, y float64
}

// NewVec2RT makes a new vec2 from the polar coordinates (r,t)
func NewVec2RT(r, t float64) Vector2 {
	return Vector2{math.Cos(t) * r, math.Sin(t) * r}
}

// Sum returns a fresh vector sum
func Sum(v1, v2 *Vector2) *Vector2 {
	return &Vector2{v1.x + v2.x, v1.y + v2.y}
}

// Minus returns a fresh vector difference
func Minus(v1, v2 *Vector2) *Vector2 {
	return &Vector2{v1.x - v2.x, v1.y - v2.y}
}

// Scale returns a freshly scaled vector
func Scale(v1 *Vector2, a float64) *Vector2 {
	return &Vector2{v1.x * a, v1.y * a}
}

// Mag returns the magnitude of the vector
func Mag(v *Vector2) float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2))
}
