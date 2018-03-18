package main

import (
	"fmt"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"

	"github.com/fogleman/gg"
)

const (
	eps = 0.0001
)

type line struct {
	o, d vec2
}

func main() {
	fmt.Println("Generating cool art...")
	radius := 400

	lines := []line{}

	fan := 1.0
	for i := 0.0; i < fan; i++ {
		spawn := line{vec2{0, 0}, newRT(rand.Float64()*30, 2*math.Pi/fan*i)}
		growLine(&spawn, &lines, radius)
	}

	for i := 0; i < 4000; i++ {
		parent := line{}
		for mag(&parent.d) < 10 {
			parent = lines[rand.Intn(len(lines))]
		}
		o := newVec2(&parent.d)
		o.times(rand.Float64()).add(&parent.o)
		d := norm(&parent.d)
		//d.div(mag(&d))
		//d.times((rand.Float64()*2 - 1) * 12)
		if rand.Intn(2) == 0 {
			d.times(-1)
		}
		spawn := line{o, d}
		growLine(&spawn, &lines, radius)

	}
	draw(lines, radius)
}

func growLine(l *line, lines *[]line, radius int) {

	o := sum(&l.o, &l.d)
	if mag(&o) > float64(radius) {
		return
	}

	t := 1.0
	for _, other := range *lines {
		t1, t2, result := intersect(&l.o, &l.d, &other.o, &other.d)
		if result == inf {
			return
		}
		if result == one {
			if t2 > eps && t2 < 1-eps && t1 > eps && t1 < 1-eps {
				if t1 < t {
					t = t1
				}
			}
		}
	}

	if t < 0.01 {
		return
	}

	l.d.times(t)
	*lines = append(*lines, *l)

	if t == 1 {
		r, th := mag(&l.d), heading(&l.d)
		d := newRT(r+rand.Float64()-0.5, th+rand.Float64()*0.4-0.2)
		spawn := line{o, d}
		growLine(&spawn, lines, radius)
	}
}

func addLine(l line, lines []line, radius int) []line {
	t := 1.0
	for _, other := range lines {
		t1, t2, result := intersect(&l.o, &l.d, &other.o, &other.d)
		if result == inf {
			return lines
		}
		if result == one {
			if t2 > eps && t2 < 1-eps && t1 > eps && t1 < 1-eps {
				if t1 < t {
					t = t1
				}
			}
		}
	}
	l.d.times(t)

	return append(lines, l)
}

func draw(lines []line, radius int) {
	margin := 10
	ctx := gg.NewContext((radius+margin)*2, (radius+margin)*2)

	ctx.SetColor(color.White)
	ctx.Clear()
	ctx.Translate(float64(radius+margin), float64(radius+margin))

	for _, l := range lines {
		ctx.MoveTo(l.o.x, l.o.y)
		ctx.LineTo(l.o.x+l.d.x, l.o.y+l.d.y)
	}
	ctx.SetLineWidth(2)
	ctx.SetColor(color.Black)
	ctx.Stroke()

	file, err := os.Create("out.png")
	if err != nil {
		fmt.Printf("Failed to create out file.\n")
	}
	png.Encode(file, ctx.Image())
}
