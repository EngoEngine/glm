package geo

import (
	"github.com/luxengine/glm"
)

// DOP8 is an 8-DOP
type DOP8 struct {
	Min [4]float32
	Max [4]float32
}

// Intersects returns true if the 8-DOP intersect.
func (d *DOP8) Intersects(o *DOP8) bool {
	for n := 0; n < 4; n++ {
		if d.Min[n] > o.Max[n] || d.Max[n] < o.Min[n] {
			return false
		}
	}
	return true
}

// ComputeFromPoints3 recomputes the 8-DOP from the given points in world space.
func (d *DOP8) ComputeFromPoints3(points []glm.Vec3) {
	// Reinitialize the 8-DOP to an empty volume.
	d.Min = [4]float32{}
	d.Max = [4]float32{}

	var value float32
	for n := 0; n < len(points); n++ {
		// Axis 0 = ( 1,  1,  1)
		value = points[n][0] + points[n][1] + points[n][2]
		if value < d.Min[0] {
			d.Min[0] = value
		} else if value > d.Max[0] {
			d.Max[0] = value
		}

		// Axis 1 = ( 1,  1, -1)
		value = points[n][0] + points[n][1] - points[n][2]
		if value < d.Min[1] {
			d.Min[1] = value
		} else if value > d.Max[1] {
			d.Max[1] = value
		}

		// Axis 2 = ( 1, -1,  1)
		value = points[n][0] - points[n][1] + points[n][2]
		if value < d.Min[2] {
			d.Min[2] = value
		} else if value > d.Max[2] {
			d.Max[2] = value
		}

		// Axis 3 = (-1,  1,  1)
		value = -points[n][0] + points[n][1] + points[n][2]
		if value < d.Min[3] {
			d.Min[3] = value
		} else if value > d.Max[3] {
			d.Max[3] = value
		}
	}
}
