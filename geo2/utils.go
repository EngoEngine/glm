package geo2

import (
	"github.com/luxengine/glm"
	"math"
)

// SqDistPointSegment returns the squared distance between point c and segment
// ab
func SqDistPointSegment(a, b, c *glm.Vec2) float32 {
	ab, ac, bc := b.Sub(a), c.Sub(a), b.Sub(c)
	e := ac.Dot(&ab)

	if e <= 0 {
		return ac.Len2()
	}
	f := ab.Len2()
	if e >= f {
		return bc.Len2()
	}

	return ac.Len2() - e*e/f
}

// MostSeparatePointsOnAABB compute indices to the two most separated points of
// the (up to) 4 points defining the AABB encompassing the point set.
func MostSeparatePointsOnAABB(points []glm.Vec2) (min, max int) {
	// First find most extreme points along principal axes
	var minx, maxx, miny, maxy int

	for i := 1; i < len(points); i++ {
		if points[i][0] < points[minx][0] {
			minx = i
		}
		if points[i][0] > points[maxx][0] {
			maxx = i
		}
		if points[i][1] < points[miny][1] {
			miny = i
		}
		if points[i][1] > points[maxy][1] {
			maxy = i
		}
	}

	// Compute the squared distances for the three pairs of points
	dx := points[maxx].Sub(&points[minx])
	dy := points[maxy].Sub(&points[miny])

	dx2 := dx.Len2()
	dy2 := dy.Len2()

	// Pick the pair (min,max) of points most distant
	min = minx
	max = maxx
	if dy2 > dx2 {
		max = maxy
		min = miny
	}
	return
}

// PointFarthestFromEdge returns the index of the point that is the farthest
// from the line ab.
func PointFarthestFromEdge(a, b *glm.Vec2, points []glm.Vec2) (index int) {
	e := b.Sub(a)
	eperp := e.Perp()

	index = -1
	var maxVal, rightMostVal float32

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

// CovarianceMatrix2 computes the covariance matrix of the given set of points.
func CovarianceMatrix2(cov *glm.Mat3, points []glm.Vec3) {
	oon := float32(1.0) / float32(len(points))
	var c glm.Vec3
	var e00, e11, e01 float32
	// Compute the center of mass (centroid) of the points
	for i := range points {
		c.AddWith(&points[i])
	}

	c.MulWith(oon)

	// Compute covariance elements
	for i := range points {
		// Translate points so center of mass is at origin
		p := points[i].Sub(&c)

		// Compute covariance of translated points
		e00 += p[0] * p[0]
		e11 += p[1] * p[1]

		e01 += p[0] * p[1]
	}

	// Fill in the covariance matrix elements
	cov[0] = e00 * oon
	cov[3] = e11 * oon

	cov[1] = e01 * oon

	cov[2] = cov[1]
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

// ClosestPointOnLine2 returns the point on ab closest to c. Also returns t for
// the position of d, d(t) = a + t*(b - a)
func ClosestPointOnLine2(a, b, c *glm.Vec2) (t float32, point glm.Vec2) {
	ab := b.Sub(a)

	// Project c onto ab, but deferring the division by ab.Dot(ab)
	cma := c.Sub(a)
	t = cma.Dot(&ab)

	if t <= 0 {
		// 'c' projects outside the [a, b] interval, on the 'a' side; clamp to
		// 'a'
		return 0, *a
	}

	denom := ab.Dot(&ab)
	if t >= denom {
		// 'c' projects outside the [a, b] interval, on the 'b' side; clamp to
		// 'b'
		return 1, *b
	}

	// 'c' projects inside the [a, b] interval; most do the deferred divide now
	t = t / denom
	point = *a
	point.AddScaledVec(t, &ab)

	return
}
