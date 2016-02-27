package debug

import (
	"fmt"
	"math"
	"os"

	"github.com/ajstarks/svgo"
	"github.com/luxengine/glm"
)

// Draw draws all the elements to a svg image
func Draw(filename string, elements ...glm.Vec2) error {
	os.Remove(filename)
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	c := svg.New(f)

	min := [2]float32{math.MaxFloat32, math.MaxFloat32}
	max := [2]float32{-math.MaxFloat32, -math.MaxFloat32}

	for n := 0; n < len(elements); n++ {
		if min[0] > elements[n][0] {
			min[0] = elements[n][0]
		}

		if min[1] > elements[n][1] {
			min[1] = elements[n][1]
		}

		if max[0] < elements[n][0] {
			max[0] = elements[n][0]
		}

		if max[1] < elements[n][1] {
			max[1] = elements[n][1]
		}
	}

	padding := 80

	fmt.Println(int(max[0]-min[0])+padding, int(max[1]-min[1])+padding)

	c.Start(int(max[0]-min[0])+padding, int(max[1]-min[1])+padding)

	for n := 0; n < len(elements); n++ {
		v1 := elements[n]
		v2 := elements[(n+1)%len(elements)]
		c.Circle(int(v1[0]-min[0])+padding/2, int(v1[1]-min[1])+padding/2, 3, "fill:black")
		c.Line(int(v1[0]-min[0])+padding/2, int(v1[1]-min[1])+padding/2, int(v2[0]-min[0])+padding/2, int(v2[1]-min[1])+padding/2, "stroke:blue;stroke-width:2")

		midpoint := glm.Vec2{(v1[0] + v2[0]) / 2, (v1[1] + v2[1]) / 2}
		e := glm.Vec2{v2[0] - v1[0], v2[1] - v1[1]}
		eperp := glm.Vec2{-e[1], e[0]}
		eperp.Normalize()
		lmid := midpoint.Len()

		c.Line(int(midpoint[0]-min[0])+padding/2, int(midpoint[1]-min[1])+padding/2,
			int(midpoint[0]+eperp[0]*lmid/8-min[0])+padding/2, int(midpoint[1]+eperp[1]*lmid/8-min[1])+padding/2,
			"stroke:green;stroke-width:2")
	}

	c.End()
	return nil
}
