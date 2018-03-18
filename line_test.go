package main

import (
	"testing"
)

func TestIntersect(t *testing.T) {
	cases := []struct {
		o1, d1, o2, d2 vec2
		t1, t2         float64
		result         int
	}{
		{vec2{0, 0}, vec2{1, 0}, vec2{3, 2}, vec2{0, -1},
			3, 2, one},
		{vec2{1, 0}, vec2{1, 1}, vec2{1, 2}, vec2{1, -1},
			1, 1, one},
		{vec2{0, 0}, vec2{1, 0}, vec2{2, 0}, vec2{-1, 0},
			0, 0, inf},
		{vec2{0, 0}, vec2{1, 0}, vec2{2, 0.1}, vec2{-1, 0},
			0, 0, zero},
	}

	for _, c := range cases {
		t1, t2, result := intersect(&c.o1, &c.d1, &c.o2, &c.d2)
		if t1 != c.t1 || t2 != c.t2 || result != c.result {
			t.Errorf("intersect(%v,%v,%v,%v) == %v,%v,%v, want %v,%v,%v",
				c.o1, c.d1, c.o2, c.d2, t1, t2, result, c.t1, c.t2, c.result)
		}
	}
}
