package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"

	"github.com/fogleman/gg"
)

type line struct {
	a Vector2
	b Vector2
}

func main() {
	lines := make([]line, 0)

	lines = append(lines, line{Vector2{0, 0}, Vector2{0, 20}})
	unvisited := []int{0}

	for i := 0; i < 1000; i++ {
		lines, unvisited = fracture(lines, unvisited, 2)
	}

	fmt.Println(len(lines))

	img := drawLines(lines)
	file, _ := os.Create("out.png")
	png.Encode(file, img)
}

func drawLines(lines []line) image.Image {
	size := 1600

	ctx := gg.NewContext(size, size)
	ctx.Translate(float64(size/2), float64(size/2))
	ctx.SetColor(color.White)
	for _, line := range lines {
		ctx.MoveTo(line.a.x, line.a.y)
		ctx.LineTo(line.b.x, line.b.y)
		ctx.Stroke()
	}
	return ctx.Image()
}

func fracture(lines []line, unvisited []int, n int) ([]line, []int) {
	maxLen := 30.0

	visited := make([]int, 0)

	for i := 0; i < n; i++ {
		index := unvisited[rand.Intn(len(unvisited))]
		parent := lines[index]
		visited = append(visited, index)

		delta := rand.Float64()
		spawn := Sum(Scale(&parent.a, delta), Scale(&parent.b, 1-delta))

		parentDir := Minus(&parent.b, &parent.a)
		dir := Vector2{parentDir.y, -parentDir.x}
		dir.x += rand.Float64()*2 - 1
		dir.y += rand.Float64()*2 - 1
		dir = *Scale(&dir, (2*rand.Float64()-1)*maxLen/Mag(&dir))

		l := line{*spawn, *Sum(spawn, &dir)}

		unvisited = append(unvisited, len(lines))
		lines = append(lines, l)
	}

	for _, index := range visited {
		for i := 0; i < len(unvisited); i++ {
			if unvisited[i] == index {
				unvisited[i] = unvisited[len(unvisited)-1]
				unvisited = unvisited[:len(unvisited)-1]
				break
			}
		}

	}

	return lines, unvisited
}
