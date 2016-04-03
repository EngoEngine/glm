package main

import (
	"github.com/ajstarks/svgo/float"
	"github.com/luxengine/glm"
	"github.com/luxengine/glm/geo"
	"os"
)

func main() {
	const (
		width  = 500
		height = 500
	)
	c := svg.New(os.Stdout)
	c.Start(500, 500)

	s := geo.Simplex2{
		Points: [3]glm.Vec2{{100, 100}, {100, 200}, {200, 100}},
		Size:   3,
	}

	for n := 0; n < 3; n++ {
		l := s.Points[(n+1)%3].Sub(&s.Points[n])
		c.Line(float64(s.Points[n][0]), float64(s.Points[n][1]), float64(s.Points[n][0]+l[0]), float64(s.Points[n][1]+l[1]), "stroke:rgb(0,0,255);stroke-width:2")
	}

	const (
		sx = 25
		sy = 25
	)
	for n := 0; n < sx; n++ {
		for m := 0; m < sy; m++ {
			p := glm.Vec2{float32(n) * float32(width) / float32(sx), float32(m) * float32(height) / float32(sy)}

			s2 := geo.Simplex2{
				Points: [3]glm.Vec2{s.Points[0].Sub(&p), s.Points[1].Sub(&p), s.Points[2].Sub(&p)},
				Size:   3,
			}

			dir, contain := s2.NearestToOrigin()

			if !dir.ApproxEqual(&glm.Vec2{}) {
				dir.Normalize()
				dir.MulWith(10)
				var col string
				if contain {
					col = "stroke:rgb(0,255,0);stroke-width:1"
				} else {
					col = "stroke:rgb(255,0,0);stroke-width:1"
				}
				c.Line(float64(p[0]), float64(p[1]), float64(p[0]+dir[0]), float64(p[1]+dir[1]), col)
				c.Circle(float64(p[0]), float64(p[1]), 2, "fill:rgb(255,0,0)")
			} else {
				c.Circle(float64(p[0]), float64(p[1]), 2, "fill:rgb(0,255,0)")
			}
		}
	}

	c.End()
}
