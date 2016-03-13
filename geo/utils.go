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
	eperp := glm.Vec2{-e[1], e[0]}

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

// MostSeparatePointsOnAABB2 compute indices to the two most separated points of
// the (up to) six points defining the AABB encompassing the point set.
func MostSeparatePointsOnAABB2(points []glm.Vec2) (min, max int) {
	// First find most extreme points along principal axes
	var minx, maxx, miny, maxy, minz, maxz int

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
	dz := points[maxz].Sub(&points[minz])

	dx2 := dx.Len2()
	dy2 := dy.Len2()
	dz2 := dz.Len2()

	// Pick the pair (min,max) of points most distant
	min = minx
	max = maxx
	if dy2 > dx2 && dy2 > dz2 {
		max = maxy
		min = miny
	}
	if dz2 > dx2 && dz2 > dy2 {
		max = maxz
		min = minz
	}
	return
}

// Variance computes the variance of a set of 1D values.
func Variance(x []float32) float32 {
	var u float32
	for i := range x {
		u += x[i]
	}
	u /= float32(len(x))
	var s2 float32
	for i := range x {
		s2 += (x[i] - u) * (x[i] - u)
	}
	return s2 / float32(len(x))
}

// CovarianceMatrix3 computes the covariance matrix of the given set of points.
func CovarianceMatrix3(cov *glm.Mat3, points []glm.Vec3) {
	oon := float32(1.0) / float32(len(points))
	var c glm.Vec3
	var e00, e11, e22, e01, e02, e12 float32
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
		e22 += p[2] * p[2]
		e01 += p[0] * p[1]
		e02 += p[0] * p[2]
		e12 += p[1] * p[2]
	}

	//     0 1 2
	//   X------
	// 0 | 0 3 6
	// 1 | 1 4 7
	// 2 | 2 5 8

	// Fill in the covariance matrix elements
	cov[0] = e00 * oon
	cov[4] = e11 * oon
	cov[8] = e22 * oon

	cov[1] = e01 * oon
	cov[2] = e02 * oon
	cov[5] = e12 * oon

	cov[3] = cov[1]
	cov[6] = cov[2]
	cov[7] = cov[5]
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

// SymSchur2 aka: 2-by-2 Symmetric Schur decomposition. Given an n-by-n symmetric matrix
// and indices p, q such that 1 <= p < q <= n, computes a sine-cosine pair
// (s, c) that will serve to form a Jacobi rotation matrix.
//
// See Golub, Van Loan, Matrix Computations, 3rd ed, p428
func SymSchur2(a *glm.Mat3, p, q int) (c, s float32) {
	if math.Abs(a[3*q+p]) > 0.0001 {
		r := (a[3*q+q] - a[3*p+p]) / (2.0 * a[3*q+p])
		var t float32
		if r >= 0 {
			t = 1.0 / (r + math.Sqrt(1.0+r*r))
		} else {
			t = -1.0 / (-r + math.Sqrt(1.0+r*r))
		}
		c = 1.0 / math.Sqrt(1.0+t*t)
		s = t * c
	} else {
		c = 1.0
		s = 0.0
	}
	return
}

// Jacobi computes the eigenvectors and eigenvalues of the symmetric matrix A
// using the classic Jacobi method of iteratively updating A as A = J∧T * A * J,
// where J = J(p, q, theta) is the Jacobi rotation matrix.
//
// On exit, v will contain the eigenvectors, and the diagonal elements
// of a are the corresponding eigenvalues.
//
// See Golub, Van Loan, Matrix Computations, 3rd ed, p428
func Jacobi(a, v *glm.Mat3) {
	const maxIterations = 50

	var i, j, n, p, q int
	var prevoff, c, s float32
	var J glm.Mat3
	// Initialize v to identify matrix
	for i = 0; i < 3; i++ {
		v[3*0+i] = 0
		v[3*1+i] = 0
		v[3*2+i] = 0
		v[3*i+i] = 1
	}

	// Repeat for some maximum number of iterations
	for n = 0; n < maxIterations; n++ {
		// Find largest off-diagonal absolute element a[p][q]
		p, q = 0, 1
		for i = 0; i < 3; i++ {
			for j = 0; j < 3; j++ {
				if i == j {
					continue
				}
				if math.Abs(a[3*j+i]) > math.Abs(a[3*q+p]) {
					p = i
					q = j
				}
			}
		}
		// Compute the Jacobi rotation matrix J(p, q, theta)
		// (This code can be optimized for the three different cases of rotation)
		c, s = SymSchur2(a, p, q)
		for i = 0; i < 3; i++ {
			J[3*0+i] = 0
			J[3*1+i] = 0
			J[3*2+i] = 0
			J[3*i+i] = 1
		}
		J[3*p+p] = c
		J[3*q+p] = s
		J[3*p+q] = -s
		J[3*q+q] = c

		// Cumulate rotations into what will contain the eigenvectors
		*v = v.Mul3(&J)
		// Make ’a’ more diagonal, until just eigenvalues remain on diagonal

		Jt := J.Transposed()
		Jta := Jt.Mul3(a)
		a.Mul3Of(&Jta, &J)

		// Compute "norm" of off-diagonal elements
		var off float32
		for i = 0; i < 3; i++ {
			for j = 0; j < 3; j++ {
				if i == j {
					continue
				}
				off += a[3*j+i] * a[3*j+i]
			}
		}
		/* off = sqrt(off); not needed for norm comparison */

		// Stop when norm no longer decreasing
		if n > 2 && off >= prevoff {
			return
		}
		prevoff = off
	}
}

// MinimumAreaRectangle returns the center point and axis orientation of the
// minimum area rectangle in the xy plane.
func MinimumAreaRectangle(points []glm.Vec2) (minArea float32, center glm.Vec2, orientation [2]glm.Vec2) {
	minArea = float32(math.MaxFloat32)

	// Loop through all edges; j trails i by 1, modulo len(points)
	for i, j := 0, len(points)-1; i < len(points); i++ {
		// Get current edge e0 (e0x, e0y), normalized
		e0 := points[i].Sub(&points[j])
		e0.Normalize()

		// Get an axis e1 orthogonal to edge e0
		e1 := glm.Vec2{-e0[1], e0[0]}

		var min0, min1, max0, max1 float32
		for k := 0; k < len(points); k++ {
			// Project points onto axes e0 and e1 and keep track of minimum and
			// maximum values along both axes.
			d := points[k].Sub(&points[j])

			dot := d.Dot(&e0)
			if dot < min0 {
				min0 = dot
			}

			if dot > max0 {
				max0 = dot
			}

			dot = d.Dot(&e1)
			if dot < min1 {
				min1 = dot
			}

			if dot > max1 {
				max1 = dot
			}
		}
		area := (max0 - min0) * (max1 - min1)

		// If best so far, remember area, center, and axes.
		if area < minArea {
			minArea = area
			orientation[0] = e0
			orientation[1] = e1

			t0 := e0.Mul(min0 + max0)
			t1 := e1.Mul(min1 + max1)
			t0.AddWith(&t1)
			t0.MulWith(0.5)

			center = points[j].Add(&t0)
		}

		// trail i
		j = i
	}
	return
}

// ClosestPointSegmentSegment computes points C₁ and C₂ of
// S₁(s) = p₁ + s * (q₁-p₁) and S₂(t) = p₂ + t * (q₂-p₂), returning 's', 't', and the
// squared distance 'u' between S₁(s) and S₂(t).
func ClosestPointSegmentSegment(p1, q1, p2, q2 *glm.Vec3) (s, t, u float32, c1, c2 glm.Vec3) {
	const (
		epsilon = 0.000
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

// SqDistPointSegment2 returns the squared distance between point c and segment
// ab
func SqDistPointSegment2(a, b, c *glm.Vec2) float32 {
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

// SqDistPointSegment3 returns the squared distance between point c and segment
// ab
func SqDistPointSegment3(a, b, c *glm.Vec3) float32 {
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
