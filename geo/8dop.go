package geo

import (
	"github.com/luxengine/glm"
)

// DOP8 is an 8-DOP.
type DOP8 struct {
	Min [4]float32
	Max [4]float32
}

// TestDOP8DOP8 returns true if the 8-DOP intersect.
func TestDOP8DOP8(a, b *DOP8) bool {
	for n := 0; n < 4; n++ {
		if a.Min[n] > b.Max[n] || a.Max[n] < b.Min[n] {
			return false
		}
	}
	return true
}

// DOP8FromPoints recomputes the 8-DOP from the given points in world space.
func DOP8FromPoints(d *DOP8, points []glm.Vec3) {
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
