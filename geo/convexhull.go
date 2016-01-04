package geo

import (
	"github.com/luxengine/glm"
	"github.com/luxengine/math"
)

// PointFarthestFromEdge returns the index of the point that is the farthest
// from the line a-b
func PointFarthestFromEdge(a, b *glm.Vec2, points []glm.Vec2) int {
	e := b.Sub(a)
	eperp := glm.Vec2{-e.Y(), e.X()}

	bestindex := -1
	maxVal := -float32(math.MaxFloat32)
	rightMostVal := -float32(math.MaxFloat32)

	for n := 0; n < len(points); n++ {
		pma := points[n].Sub(a)
		d := pma.Dot(&eperp)
		r := pma.Dot(&e)
		if d > maxVal || (d == maxVal && r > rightMostVal) {
			maxVal = d
			bestindex = n
			rightMostVal = r
		}
	}

	return bestindex
}
