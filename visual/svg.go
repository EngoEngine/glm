package visual

import (
	"fmt"
	"github.com/ajstarks/svgo/float"
	"github.com/luxengine/glm/geo2"
	"os"
)

const (
	picW   = 500
	picH   = 500
	buffer = 2
)

type svgw struct {
	startX, startY, width, height float32
	s                             *svg.SVG
}

func (s svgw) Line(x0, y0, x1, y1 float32) {
	s.s.Line(
		float64(((x0-s.startX)/s.width)*picW),
		picH-float64(((y0-s.startY)/s.height)*picH),
		float64(((x1-s.startX)/s.width)*picW),
		picH-float64(((y1-s.startY)/s.height)*picH),
		"stroke:black;stroke-width:2;",
	)
}

func (s svgw) Point(x, y float32) {
	s.s.Circle(float64(((x-s.startX)/s.width)*picW),
		picH-float64(((y-s.startY)/s.height)*picH),
		3, "fill:black")
}

// Visualise takes a bunch of elements and renders them in a svg.
func Visualise(filename string, args ...interface{}) {
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	var maxsize [4]float32
	for _, arg := range args {
		a := size(arg)
		eat(&maxsize, &a)
	}

	s := svgw{
		s:      svg.New(file),
		startX: maxsize[0] - buffer,
		startY: maxsize[1] - buffer,
		width:  maxsize[2] - maxsize[0] + buffer*2,
		height: maxsize[3] - maxsize[1] + buffer*2,
	}

	s.s.Start(picW, picH)

	for _, arg := range args {
		draw(s, arg)
	}

	s.s.End()
}

func eat(fill, food *[4]float32) {
	if food[0] < fill[0] {
		fill[0] = food[0]
	}

	if food[1] < fill[1] {
		fill[1] = food[1]
	}

	if food[2] > fill[2] {
		fill[2] = food[2]
	}

	if food[3] > fill[3] {
		fill[3] = food[3]
	}
}

// return [minx, miny, maxx, maxy]
func size(arg interface{}) [4]float32 {
	switch t := arg.(type) {
	case geo2.AABB:
		return [4]float32{
			t.Center[0] - t.HalfExtend[0],
			t.Center[1] - t.HalfExtend[1],
			t.Center[0] + t.HalfExtend[0],
			t.Center[1] + t.HalfExtend[1],
		}
	}
	return [4]float32{}
}

func draw(s svgw, arg interface{}) {
	switch t := arg.(type) {
	case geo2.AABB:
		low := [2]float32{t.Center[0] - t.HalfExtend[0],
			t.Center[1] - t.HalfExtend[1]}
		high := [2]float32{t.Center[0] + t.HalfExtend[0],
			t.Center[1] + t.HalfExtend[1]}
		s.Line(low[0], low[1], low[0], high[1])
		s.Line(low[0], high[1], high[0], high[1])
		s.Line(low[0], low[1], high[0], low[1])
		s.Line(high[0], low[1], high[0], high[1])
	}
}
