package main

import "math"

type vec2 struct {
	x, y float64
}

// NewRT makes a polar vec2
func newRT(r, t float64) vec2 {
	return vec2{r * math.Cos(t), r * math.Sin(t)}
}

func newVec2(v *vec2) vec2 {
	return vec2{v.x, v.y}
}

func sum(v1, v2 *vec2) vec2 {
	return vec2{v1.x + v2.x, v1.y + v2.y}
}

func sub(v1, v2 *vec2) vec2 {
	return vec2{v1.x - v2.x, v1.y - v2.y}
}

func neg(v *vec2) vec2 {
	return vec2{-v.x, -v.y}
}

func norm(v *vec2) vec2 {
	return vec2{-v.y, v.x}
}

func lerp(v1, v2 *vec2, delta float64) vec2 {
	return vec2{
		v1.x*(1-delta) + v2.x*delta,
		v1.y*(1-delta) + v2.y*delta,
	}
}

// Methods

func (v *vec2) set(v2 *vec2) *vec2 {
	v.x = v2.x
	v.y = v2.y
	return v
}

func (v *vec2) add(v2 *vec2) *vec2 {
	v.x += v2.x
	v.y += v2.y
	return v
}

func (v *vec2) sub(v2 *vec2) *vec2 {
	v.x -= v2.x
	v.y -= v2.y
	return v
}

func (v *vec2) times(a float64) *vec2 {
	v.x *= a
	v.y *= a
	return v
}

func (v *vec2) div(a float64) *vec2 {
	v.x /= a
	v.y /= a
	return v
}

func mag(v *vec2) float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2))
}

func heading(v *vec2) float64 {
	return math.Atan2(v.y, v.x)
}

func dot(v1, v2 *vec2) float64 {
	return v1.x*v2.x + v1.y*v2.y
}

func cross(v1, v2 *vec2) float64 {
	return v1.x*v2.y - v2.x*v1.y
}
