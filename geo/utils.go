package geo

import (
	"github.com/luxengine/glm"
	"github.com/luxengine/math"
)

// IsConvexQuad returns true if the qualidrateral is convex.
func IsConvexQuad(a, b, c, d *glm.Vec3) bool {
	dmb, amb, cmb := d.Sub(b), a.Sub(b), c.Sub(b)
	bda, bdc := dmb.Cross(&amb), dmb.Cross(&cmb)

	if bda.Dot(&bdc) >= 0 {
		return false
	}

	cma, dma, bma := c.Sub(a), d.Sub(a), b.Sub(a)
	acd := cma.Cross(&dma)
	acb := cma.Cross(&bma)
	return acd.Dot(&acb) < 0
}

// PointFarthestFromEdge returns the index of the point that is the farthest
// from the line a-b.
func PointFarthestFromEdge(a, b *glm.Vec2, points []glm.Vec2) (index int) {
	e := b.Sub(a)
	eperp := glm.Vec2{-e.Y(), e.X()}

	index = -1
	maxVal := float32(0)
	rightMostVal := float32(0)

	for n := 0; n < len(points); n++ {
		pma := points[n].Sub(a)
		d := pma.Dot(&eperp)
		r := pma.Dot(&e)
		if d > maxVal || (d == maxVal && r > rightMostVal) {
			maxVal = d
			index = n
			rightMostVal = r
		}
	}

	return
}

// ExtremePointsAlongDirection2 returns indices imin and imax into points of the
// least and most, respectively, distant points along the direction dir.
func ExtremePointsAlongDirection2(direction *glm.Vec2, points []glm.Vec2) (imin int, imax int) {

	imin, imax = -1, -1

	var minproj, maxproj float32 = math.MaxFloat32, -math.MaxFloat32

	for n := 0; n < len(points); n++ {

		// project this point along the direction
		proj := points[n].Dot(direction)

		// keep track of the least distant point along the direction vector
		if proj < minproj {
			minproj = proj
			imin = n
		}

		// keep track of the most distant point along the direction vector
		if proj > maxproj {
			maxproj = proj
			imax = n
		}
	}
	return
}

// ExtremePointsAlongDirection3 returns indices imin and imax into points of the
// least and most, respectively, distant points along the direction dir.
func ExtremePointsAlongDirection3(direction *glm.Vec3, points []glm.Vec3) (imin int, imax int) {

	imin, imax = -1, -1

	var minproj, maxproj float32 = math.MaxFloat32, -math.MaxFloat32

	for n := 0; n < len(points); n++ {

		// project this point along the direction
		proj := points[n].Dot(direction)

		// keep track of the least distant point along the direction vector
		if proj < minproj {
			minproj = proj
			imin = n
		}

		// keep track of the most distant point along the direction vector
		if proj > maxproj {
			maxproj = proj
			imax = n
		}
	}
	return
}
