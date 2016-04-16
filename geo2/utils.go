package geo2

import (
	"github.com/luxengine/glm"
	"github.com/luxengine/glm/flops/32/flops"
	"github.com/luxengine/math"
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

// ClosestPointSegmentPoint returns the point on ab closest to c. Also returns t for
// the position of d, d(t) = a + t*(b - a)
func ClosestPointSegmentPoint(a, b, c *glm.Vec2) (t float32, point glm.Vec2) {
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

// ClosestPointTrianglePoint returns the point on the triangle abc that is closest
// to p
func ClosestPointTrianglePoint(p, a, b, c *glm.Vec2) glm.Vec2 {
	ab, ac, ap := b.Sub(a), c.Sub(a), p.Sub(a)

	// Check if P in vertex region outside A
	d1, d2 := ab.Dot(&ap), ac.Dot(&ap)
	if d1 <= 0 && d2 <= 0 {
		return *a // barycentric coordinates (1, 0, 0)
	}

	bp := p.Sub(b)
	d3, d4 := ab.Dot(&bp), ac.Dot(&ap)
	if d3 >= 0 && d4 <= d3 {
		return *b // barycentric coordinates (0, 1, 0)
	}

	// Check if P in edge region of AB, if so return projection of P onto AB
	vc := d1*d4 - d3*d2
	if vc <= 0 && d1 >= 0 && d3 <= 0 {
		ret := *a
		ret.AddScaledVec(d1/(d1-d3), &ab)
		return ret
	}

	// Check if P in vertex region outside C
	cp := p.Sub(c)
	d5, d6 := ab.Dot(&cp), ac.Dot(&cp)
	if d6 >= 0 && d5 <= d6 {
		return *c // barycentric coordinates (0, 0, 1)
	}

	vb := d5*d2 - d1*d6
	if vb <= 0 && d2 >= 0 && d6 <= 0 {
		ret := *a
		ret.AddScaledVec(d2/(d2-d6), &ac)
		return ret
	}

	// Check if P in edge region of BC, if so return projection of P onto BC
	va := d3*d6 - d5*d4
	if va <= 0 && (d4-d3) >= 0 && (d5-d6) >= 0 {
		bc := c.Sub(b)
		ret := *b
		ret.AddScaledVec((d4-d3)/((d4-d3)+(d5-d6)), &bc)
		return ret // barycentric coordinates (0, 1-w, w)
	}

	// P inside face region. Compute Q through it's barycentric coordinates
	denom := 1 / (va + vb + vc)
	v := vb * denom
	w := vc * denom
	ret := *a
	ret.AddScaledVec(v, &ab)
	ret.AddScaledVec(w, &ac)
	return ret
}

// ClosestPointSegmentSegment computes points C₁ and C₂ of
// S₁(s) = p₁ + s * (q₁-p₁) and S₂(t) = p₂ + t * (q₂-p₂), returning s, t, and the
// squared distance u between S₁(s) and S₂(t).
func ClosestPointSegmentSegment(p1, q1, p2, q2 *glm.Vec2) (s, t, u float32, c1, c2 glm.Vec2) {
	// TODO(hydroflame): find a good constant for that epsilon
	const (
		epsilon = 0.0001
	)

	d1 := q1.Sub(p1)
	d2 := q2.Sub(p2)
	r := p1.Sub(p2)
	a, e, f := d1.Len2(), d2.Len2(), d2.Dot(&r)

	// Check if either or both segments degenerate into points
	if a <= epsilon && e <= epsilon {
		return 0, 0, r.Len2(), *p1, *p2
	}

	if a <= epsilon {
		// First segment degenerates into a point.
		s = 0
		t = f / e
		t = math.Clamp(t, 0, 1)
	} else {
		c := d1.Dot(&r)
		if e <= epsilon {
			// Second segment denegerates into a point.
			t = 0
			s = math.Clamp(-c/a, 0, 1)
		} else {
			// The general non-degenerate case starts here
			b := d1.Dot(&d2)
			denom := a*e - b*b // Always positive

			// If segments are not parallel, compute closest point on L₁ to L₂
			// and clamp to segment S₁. Else pick arbitrary 's' (here 0)
			if denom != 0 {
				s = math.Clamp((b*f-c*e)/denom, 0, 1)
			} else {
				s = 0
			}

			t = (b*s + f) / e

			if t < 0 {
				t = 0
				s = math.Clamp(-c/a, 0, 1)
			} else {
				t = 1
				s = math.Clamp((b-c)/a, 0, 1)
			}
		}
	}

	c1 = *p1
	c2 = *p2

	c1.AddScaledVec(s, &d1)
	c2.AddScaledVec(s, &d2)

	c1mc2 := c1.Sub(&c2)

	u = c1mc2.Len2()

	return
}

// TestSegmentSegment tests if segments ab and cd overlap. If they do, compute
// and return intersection t value along ab and intersection position p.
func TestSegmentSegment(a, b, c, d *glm.Vec2) (t float32, v glm.Vec2, overlap bool) {
	// Sign of areas correspond to which side of ab points c and d are
	a1 := Signed2DTriArea(a, b, d) // Compute winding of abd (+ or -)
	a2 := Signed2DTriArea(a, b, c) // To intersect, must have sign opposite of a1
	// If c and d are on different sides of ab, areas have different signs
	if !flops.Z(a1) && !flops.Z(a2) && a1*a2 < 0 {
		// Compute signs for a and b with respect to segment cd
		a3 := Signed2DTriArea(c, d, a) // Compute winding of cda (+ or -)
		// Since area is constant a1 - a2 = a3 - a4, or a4 = a3 + a2 - a1
		a4 := a3 + a2 - a1
		// Points a and b on different sides of cd if areas have different signs
		if !flops.Z(a3) && !flops.Z(a4) && a3*a4 < 0 {
			// Segments intersect. Find intersection point along L(t) = a + t * (b - a).
			// Given height h1 of an over cd and height h2 of b over cd,
			// t = h1 / (h1 - h2) = (b*h1/2) / (b*h1/2 - b*h2/2) = a3 / (a3 - a4),
			// where b (the base of the triangles cda and cdb, i.e., the length
			// of cd) cancels out.
			t = a3 / (a3 - a4)
			tmp := b.Sub(a)
			tmp.MulWith(t)
			v.AddOf(a, &tmp)
			overlap = true
			return
		}
	}
	// Segments not intersecting (or collinear)
	return
}

// Signed2DTriArea returns 2 times the signed triangle area. The result is
// positive if abc is CCW, negative if abc is CW, zero if abc is degenerate.
func Signed2DTriArea(a, b, c *glm.Vec2) float32 {
	return (a[0]-c[0])*(b[1]-c[1]) - (a[1]-c[1])*(b[0]-c[0])
}
