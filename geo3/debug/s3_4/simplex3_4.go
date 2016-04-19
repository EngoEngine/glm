package main

import (
	"fmt"
	"github.com/luxengine/glm"
	"github.com/luxengine/glm/geo3"
)

const num = 4

func main() {
	s := geo3.Simplex{
		Points: [4]glm.Vec3{
			{0, 0, 0},
			{0, 1, 0},
			{1, 0, 0},
			{0, 0, 1},
		},
		Size: 4,
	}

	for x := 0; x < num; x++ {
		for y := 0; y < num; y++ {
			for z := 0; z < num; z++ {
				p := glm.Vec3{
					(float32(x)/num)*2 - 0.5,
					(float32(y)/num)*2 - 0.5,
					(float32(z)/num)*2 - 0.5,
				}

				s2 := geo3.Simplex{
					Points: [4]glm.Vec3{s.Points[0].Sub(&p), s.Points[1].Sub(&p), s.Points[2].Sub(&p), s.Points[3].Sub(&p)},
					Size:   4,
				}
				fmt.Print(p)
				s2.NearestToOrigin()
				fmt.Println()
			}
		}
	}
}
