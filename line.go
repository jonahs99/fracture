package main

import "fmt"

const (
	one = iota
	zero
	inf
)

// intersect calculates the intersection of two lines
func intersect(o1, d1, o2, d2 *vec2) (float64, float64, int) {
	o := sub(o2, o1)
	nd2 := neg(d2)

	det := cross(d1, nd2)

	fmt.Printf("o: %v\n", o)
	fmt.Printf("det: %v\n", det)

	if det == 0 {
		if cross(o, d1) == 0 {
			return 0, 0, inf
		}
		return 0, 0, zero
	}

	t1 := (nd2.y*o.x - nd2.x*o.y) / det
	t2 := (d1.x*o.y - d1.y*o.x) / det
	return t1, t2, one
}
