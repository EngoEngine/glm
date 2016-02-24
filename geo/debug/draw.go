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
		if min[0] > elements[n].X() {
			min[0] = elements[n].X()
		}

		if min[1] > elements[n].Y() {
			min[1] = elements[n].Y()
		}

		if max[0] < elements[n].X() {
			max[0] = elements[n].X()
		}

		if max[1] < elements[n].Y() {
			max[1] = elements[n].Y()
		}
	}

	padding := 80

	fmt.Println(int(max[0]-min[0])+padding, int(max[1]-min[1])+padding)

	c.Start(int(max[0]-min[0])+padding, int(max[1]-min[1])+padding)

	for n := 0; n < len(elements); n++ {
		v1 := elements[n]
		v2 := elements[(n+1)%len(elements)]
		c.Circle(int(v1.X()-min[0])+padding/2, int(v1.Y()-min[1])+padding/2, 3, "fill:black")
		c.Line(int(v1.X()-min[0])+padding/2, int(v1.Y()-min[1])+padding/2, int(v2.X()-min[0])+padding/2, int(v2.Y()-min[1])+padding/2, "stroke:blue;stroke-width:2")

		midpoint := glm.Vec2{(v1.X() + v2.X()) / 2, (v1.Y() + v2.Y()) / 2}
		e := glm.Vec2{v2.X() - v1.X(), v2.Y() - v1.Y()}
		eperp := glm.Vec2{-e.Y(), e.X()}
		eperp.Normalize()
		lmid := midpoint.Len()

		c.Line(int(midpoint.X()-min[0])+padding/2, int(midpoint.Y()-min[1])+padding/2,
			int(midpoint.X()+eperp.X()*lmid/8-min[0])+padding/2, int(midpoint.Y()+eperp.Y()*lmid/8-min[1])+padding/2,
			"stroke:green;stroke-width:2")
	}

	c.End()
	return nil
}
