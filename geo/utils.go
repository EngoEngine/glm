package geo

import (
	"github.com/EngoEngine/math"
	"github.com/engoengine/glm"
	"github.com/engoengine/glm/flops/32/flops"
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

// ExtremePointsAlongDirection returns indices imin and imax into points of the
// least and most, respectively, distant points along the direction dir.
func ExtremePointsAlongDirection(direction *glm.Vec3, points []glm.Vec3) (imin int, imax int) {

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

// Variance computes the variance of a float slice.
func Variance(s []float32) float32 {
	ool := 1.0 / float32(len(s))
	var u float32
	for i := range s {
		u += s[i]
	}
	u *= ool
	var s2 float32
	for i := range s {
		s2 += (s[i] - u) * (s[i] - u)
	}
	return s2 * ool
}

// CovarianceMatrix computes the covariance matrix of the given set of points.
func CovarianceMatrix(cov *glm.Mat3, points []glm.Vec3) {
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
	// TODO(hydroflame): find a good value for that
	const maxIterations = 50

	var i, j, n, p, q int
	var prevoff, c, s float32
	// Initialize v to identity matrix
	v.Ident()

	var J glm.Mat3
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
		// Make a more diagonal, until just eigenvalues remain on diagonal

		Jt := J.Transposed()
		Jta := Jt.Mul3(a)
		a.Mul3Of(&Jta, &J)

		// Compute norm of off-diagonal elements
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
// S₁(s) = p₁ + s * (q₁-p₁) and S₂(t) = p₂ + t * (q₂-p₂), returning s, t, and the
// squared distance u between S₁(s) and S₂(t).
func ClosestPointSegmentSegment(p1, q1, p2, q2 *glm.Vec3) (s, t, u float32, c1, c2 glm.Vec3) {
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

// SqDistPointSegment returns the squared distance between point c and segment
// ab
func SqDistPointSegment(a, b, c *glm.Vec3) float32 {
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

// ClosestPointSegmentPoint returns the point on ab closest to c. Also returns t for
// the position of d, d(t) = a + t*(b - a)
func ClosestPointSegmentPoint(a, b, c *glm.Vec3) (t float32, point glm.Vec3) {
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

// ClosestPointRect is a shortcut for Rect3.ClosestPoint where the rectangle is
// defined by the span of [ab, ac].
func ClosestPointRect(p, a, b, c *glm.Vec3) glm.Vec3 {
	ab := b.Sub(a)
	ac := c.Sub(a)
	d := p.Sub(a)

	// Start result at top-left corner of rect; make steps from there
	closestPoint := *a

	// Clamp p' (projection of p to plane of r) to rectangle in the across
	// direction
	dist := d.Dot(&ab)
	maxDist := ab.Len2()

	if dist >= maxDist {
		closestPoint.AddWith(&ab)
	} else if dist > 0 {
		closestPoint.AddScaledVec(dist/maxDist, &ab)
	}

	// Clamp p' to rectangle in the down direction
	dist = d.Dot(&ac)
	maxDist = ac.Len2()

	if dist >= maxDist {
		closestPoint.AddWith(&ac)
	} else if dist > 0 {
		closestPoint.AddScaledVec(dist/maxDist, &ac)
	}

	return closestPoint
}

// ClosestPointTrianglePoint returns the point on the triangle abc that is closest
// to p
func ClosestPointTrianglePoint(p, a, b, c *glm.Vec3) glm.Vec3 {
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

// PointOutsidePlane returns true if p is outside or on triangle abc CCW.
func PointOutsidePlane(p, a, b, c *glm.Vec3) bool {
	//return Dot(p-a,Cross(b-a,c-a))>=0.0f; //[APABAC]>=0
	ap, ab, ac := p.Sub(a), b.Sub(a), c.Sub(a)
	abac := ab.Cross(&ac)
	d := ap.Dot(&abac)
	return d > 0 || flops.Z(d)
}

// PointsOnOppositeSideOfPlane returns true if point p is opposite of d, such that it
// doesn't matter if abc is CW or CCW
func PointsOnOppositeSideOfPlane(p1, a, b, c, p2 *glm.Vec3) bool {
	ap := p1.Sub(a)
	ad := p2.Sub(a)
	ab := b.Sub(a)
	ac := c.Sub(a)

	abac := ab.Cross(&ac)

	signp := ap.Dot(&abac)
	signd := ad.Dot(&abac)

	return signp*signd < 0
}

// ClosestPointTetrahedronPoint returns the closes point in or on tetrahedron abcd.
func ClosestPointTetrahedronPoint(p, a, b, c, d *glm.Vec3) glm.Vec3 {
	// Start out assuming point inside all halfspaces, so closest to itself
	closestPoint := *p
	var bestSqDist float32 = math.MaxFloat32

	if PointsOnOppositeSideOfPlane(p, a, b, c, d) {
		q := ClosestPointTrianglePoint(p, a, b, c)
		pq := q.Sub(p)
		sqDist := pq.Len2()
		if sqDist < bestSqDist {
			bestSqDist = sqDist
			closestPoint = q
		}
	}

	if PointsOnOppositeSideOfPlane(p, a, c, d, b) {
		q := ClosestPointTrianglePoint(p, a, c, d)
		pq := q.Sub(p)
		sqDist := pq.Len2()
		if sqDist < bestSqDist {
			bestSqDist = sqDist
			closestPoint = q
		}
	}

	if PointsOnOppositeSideOfPlane(p, a, d, b, c) {
		q := ClosestPointTrianglePoint(p, a, d, b)
		pq := q.Sub(p)
		sqDist := pq.Len2()
		if sqDist < bestSqDist {
			bestSqDist = sqDist
			closestPoint = q
		}
	}

	if PointsOnOppositeSideOfPlane(p, b, d, c, a) {
		q := ClosestPointTrianglePoint(p, b, d, c)
		pq := q.Sub(p)
		sqDist := pq.Len2()
		if sqDist < bestSqDist {
			// doesn't matter at this point
			//bestSqDist = sqDist
			closestPoint = q
		}
	}
	return closestPoint
}

// TriangleAreaFromLengths returns the area of a triangle defined by the given
// lengths. Returns NaN if the triangle does not exist.
func TriangleAreaFromLengths(a, b, c float32) float32 {
	po2 := (a + b + c) / 2
	return math.Sqrt(po2 * (po2 - a) * (po2 - b) * (po2 - c))
}

// DistToTriangle returns the distance of p to triangle {a b c}, CCW order
func DistToTriangle(p, a, b, c *glm.Vec3) float32 {
	l1, l2, l3 := b.Sub(a), c.Sub(a), p.Sub(a)
	cross := l2.Cross(&l1)
	cross.Normalize()
	return cross.Dot(&l3)
}

// ClosestPointLineTriangle returns the pair of points that are the closest from
// the line and the triangle.
func ClosestPointLineTriangle(p, q, a, b, c *glm.Vec3) (u, v glm.Vec3) {
	// TODO(hydroflame): implement
	panic("not implemented")
}

// ClosestPointTriangleTriangle returns the pair of points that are the closest
// from the triangle pair.
func ClosestPointTriangleTriangle(a, b, c, d, e, f *glm.Vec3) (u, v glm.Vec3) {
	// TODO(hydroflame): implement
	panic("not implemented")
}

// TestSpherePlane returns true if s and p intersect. The plane
// normal must be normalized.
func TestSpherePlane(s *Sphere, p *Plane) bool {
	// calculate the new center of the sphere as if the plane passed by [0 0]
	c := s.Center.Sub(&p.P)
	d := math.Abs(c.Dot(&p.N))
	return flops.Le(d, s.Radius)
}

// InsideSpherePlane returns true if s is completely inside plane p.
func InsideSpherePlane(s *Sphere, p *Plane) bool {
	c := s.Center.Sub(&p.P)
	d := c.Dot(&p.N)
	return flops.Lt(d, -s.Radius)
}

// TestSphereHalfspace returns true if s is touching or inside halfspace p.
func TestSphereHalfspace(s *Sphere, p *Plane) bool {
	c := s.Center.Sub(&p.P)
	d := c.Dot(&p.N)
	return flops.Le(d, s.Radius)
}

// TestOBBPlane returns true if b and p intersect.
func TestOBBPlane(b *OBB, p *Plane) bool {
	// Compute the projection interval radius of b onto L(t) = b.c + t * p.n
	r := b.HalfExtend[0]*math.Abs(p.N.Dot(&b.Orientation[0])) +
		b.HalfExtend[1]*math.Abs(p.N.Dot(&b.Orientation[1])) +
		b.HalfExtend[2]*math.Abs(p.N.Dot(&b.Orientation[2]))
	// Compute distance of box center from plane
	//float s = Dot(p.n, b.c) - p.d
	c := b.Center.Sub(&p.P)
	d := c.Dot(&p.N)
	// Intersection occurs when distance s falls within [-r,+r] interval
	return math.Abs(d) <= r
}

// TestAABBPlane tests if AABB b intersects plane p.
func TestAABBPlane(b *AABB, p *Plane) bool {
	// These two lines not necessary with a (center, extents) AABB representation
	// Compute the projection interval radius of b onto L(t) = b.c + t * p.n
	r := b.HalfExtend[0]*math.Abs(p.N[0]) +
		b.HalfExtend[1]*math.Abs(p.N[1]) +
		b.HalfExtend[2]*math.Abs(p.N[2])
	// Compute distance of box center from plane
	c := b.Center.Sub(&p.P)
	s := c.Dot(&p.N)
	// Intersection occurs when distance s falls within [-r,+r] interval
	return math.Abs(s) <= r
}

// TestSphereAABB returns true if sphere s intersects AABB b
func TestSphereAABB(s *Sphere, b *AABB) bool {
	return SqDistAABBPoint(b, &s.Center) <= s.Radius*s.Radius
}

// TestSphereOBB returns true if sphere s intersects OBB b, false otherwise.
// The point p on the OBB closest to the sphere center is also returned
func TestSphereOBB(s *Sphere, b *OBB) bool {
	// Find point p on OBB closest to sphere center
	p := ClosestPointOBBPoint(b, &s.Center)
	// Sphere and OBB intersect if the (squared) distance from sphere
	// center to point p is less than the (squared) sphere radius
	v := p.Sub(&s.Center)
	return v.Dot(&v) <= s.Radius*s.Radius
}

// TestSphereTriangle returns true if sphere s intersects triangle ABC, false
// otherwise. The point p on abc closest to the sphere center is also returned.
func TestSphereTriangle(s *Sphere, a, b, c *glm.Vec3) bool {
	// Find point P on triangle ABC closest to sphere center
	p := ClosestPointTrianglePoint(&s.Center, a, b, c)
	// Sphere and triangle intersect if the (squared) distance from sphere
	// center to point p is less than the (squared) sphere radius
	v := p.Sub(&s.Center)
	return v.Dot(&v) <= s.Radius*s.Radius
}

// TestTriangleAABB returns true if [v0 v1 v2] intersects b
func TestTriangleAABB(v0, v1, v2 *glm.Vec3, b *AABB) bool {
	panic("implementation incomplete")
	/*var p0, p2, r float32
	//var p1 float32
	// Translate triangle as conceptually moving AABB to origin
	u0 := v0.Sub(&b.Center)
	u1 := v1.Sub(&b.Center)
	u2 := v2.Sub(&b.Center)
	// Compute edge vectors for triangle
	f0 := u1.Sub(&u0)
	f1 := u2.Sub(&u1)
	//f2 := u0.Sub(&u2)

	// Test axes a00..a22 (category 3)
	// Test axis a00
	p0 = u0[2]*u1[1] - u0[1]*u1[2]
	p2 = u2[2]*(u1[1]-u0[1]) - u2[2]*(u1[2]-u0[2])
	r = b.HalfExtend[1]*math.Abs(f0[2]) + b.HalfExtend[2]*math.Abs(f0[1])
	if math.Max(-math.Max(p0, p2), math.Min(p0, p2)) > r {
		return false // Axis is a separating axis
	}
	// Repeat similar tests for remaining axes a01..a22
	//...
	// Test the three axes corresponding to the face normals of AABB b (category 1).
	// Exit if...
	// ... [-b.HalfExtend[0], b.HalfExtend[0]] and [min(u0[0],u1[0],u2[0]), max(u0[0],u1[0],u2[0])] do not overlap
	if math.Max(u0[0], math.Max(u1[0], u2[0])) < -b.HalfExtend[0] ||
		math.Min(u0[0], math.Min(u1[0], u2[0])) > b.HalfExtend[0] {
		return false
	}
	// ... [-b.HalfExtend[1], b.HalfExtend[1]] and [min(u0[1],u1[1],u2[1]), max(u0[1],u1[1],u2[1])] do not overlap
	if math.Max(u0[1], math.Max(u1[1], u2[1])) < -b.HalfExtend[1] ||
		math.Min(u0[1], math.Min(u1[1], u2[1])) > b.HalfExtend[1] {
		return false
	}
	// ... [-b.HalfExtend[2], b.HalfExtend[2]] and [min(u0[2],u1[2],u2[2]), max(u0[2],u1[2],u2[2])] do not overlap
	if math.Max(u0[2], math.Max(u1[2], u2[2])) < -b.HalfExtend[2] ||
		math.Min(u0[2], math.Min(u1[2], u2[2])) > b.HalfExtend[2] {
		return false
	}
	// Test separating axis corresponding to triangle face normal (category 2)
	var p Plane
	p.N = f0.Cross(&f1)
	p.P = u0
	return TestAABBPlane(b, &p)*/
}

// IntersectSegmentPlane returns how far in the segment, the point in world
// coordinates and true if the segment and the plane intersect
func IntersectSegmentPlane(a, b *glm.Vec3, p *Plane) (t float32, q glm.Vec3, overlap bool) {
	// Compute the t value for the directed line ab intersecting the plane
	ab := glm.Vec3{
		(b[0] - p.P[0]) - (a[0] - p.P[0]),
		(b[1] - p.P[1]) - (a[1] - p.P[1]),
		(b[2] - p.P[2]) - (a[2] - p.P[2]),
	}
	t = p.N.Dot(a) / p.N.Dot(&ab)
	// If t in [0..1] compute and return intersection point
	if flops.Gez(t) && flops.Le(t, 1) {
		q = *a
		q.AddScaledVec(t, &ab)
		overlap = true
		return
	}
	// Else no intersection
	return
}

// IntersectRaySphere intersects ray r = p + td, |d| = 1, with sphere s and, if intersecting, // returns t value of intersection and intersection point q
func IntersectRaySphere(p, d *glm.Vec3, s *Sphere) (t float32, q glm.Vec3, overlap bool) {
	m := p.Sub(&s.Center)
	b := m.Dot(d)
	c := m.Dot(&m) - s.Radius*s.Radius
	// Exit if r’s origin outside s (c > 0) and r pointing away from s (b > 0)
	if flops.Gtz(c) && flops.Gtz(b) {
		return
	}
	discr := b*b - c
	// A negative discriminant corresponds to ray missing sphere
	if flops.Ltz(discr) {
		return // returns false and all zero value
	}
	// Ray now found to intersect sphere, compute smallest t value of intersection t = -b - Sqrt(discr);
	// If t is negative, ray started inside sphere so clamp t to zero
	if t < 0 {
		t = 0
	}
	q = *p
	q.AddScaledVec(t, d)
	overlap = true
	return
}

// TestRaySphere tests if ray r = p + td intersects sphere s
func TestRaySphere(p, d *glm.Vec3, s *Sphere) bool {
	m := p.Sub(&s.Center)
	c := m.Dot(&m) - s.Radius*s.Radius
	// If there is definitely at least one real root, there must be an intersection
	if flops.Lez(c) {
		return true
	}
	b := m.Dot(d)
	// Early exit if ray origin outside sphere and ray pointing away from sphere
	if flops.Gtz(b) {
		return false
	}
	disc := b*b - c
	// A negative discriminant corresponds to ray missing sphere
	if flops.Ltz(disc) {
		return false
	}
	// Now ray must hit sphere
	return true
}

// IntersectRayAABB intersect ray R(t) = p + t*d against AABB a. When
// intersecting, return intersection distance t and point q of intersection.
func IntersectRayAABB(p, d *glm.Vec3, a *AABB) (t float32, q glm.Vec3, overlap bool) {
	// TODO(hydroflame): find what epsilon to use.
	const epsilon = 0.00001
	tmax := float32(math.MaxFloat32) // set to max distance ray can travel (for segment)
	// For all three slabs
	for i := 0; i < 3; i++ {
		if math.Abs(d[i]) < epsilon {
			// Ray is parallel to slab. No hit if origin not within slab
			if p[i] < a.Center[i]-a.HalfExtend[i] || p[i] > a.Center[i]+a.HalfExtend[i] {
				return
			}
		} else {
			// Compute intersection t value of ray with near and far plane of slab
			ood := 1.0 / d[i]
			t1 := (a.Center[i] - a.HalfExtend[i] - p[i]) * ood
			t2 := (a.Center[i] + a.HalfExtend[i] - p[i]) * ood
			// Make t1 be intersection with near plane, t2 with far plane
			if t1 > t2 {
				t1, t2 = t2, t1
			}
			// Compute the intersection of slab intersection intervals
			if t1 > t {
				t = t1
			}
			if t2 > tmax {
				tmax = t2
			}
			// Exit with no collision as soon as slab intersection becomes empty
			if t > tmax {
				return
			}

		}
	}
	// Ray intersects all 3 slabs. Return point (q) and intersection t value (tmin)
	q = *p
	q.AddScaledVec(t, d)
	overlap = true
	return
}

// TestSegmentAABB tests if segment specified by points p0 and p1 intersects
// AABB b.
func TestSegmentAABB(p0, p1 *glm.Vec3, b *AABB) bool {
	// TODO(hydroflame): find what epsilon to use.
	const epsilon = 0.00001

	m := p0.Add(p1) // Segment midpoint
	m.MulWith(0.5)
	d := p1.Sub(&m)      // Segment halflength vector
	m = m.Sub(&b.Center) // Translate box and segment to origin

	// Try world coordinate axes as separating axes
	adx := math.Abs(d[0])
	if flops.Gt(math.Abs(m[0]), b.HalfExtend[0]+adx) {
		return false
	}
	ady := math.Abs(d[1])
	if flops.Gt(math.Abs(m[1]), b.HalfExtend[1]+ady) {
		return false
	}
	adz := math.Abs(d[2])
	if flops.Gt(math.Abs(m[2]), b.HalfExtend[2]+adz) {
		return false
	}
	// Add in an epsilon term to counteract arithmetic errors when segment is
	// (near) parallel to a coordinate axis (see text for detail)
	adx += epsilon
	ady += epsilon
	adz += epsilon
	// Try cross products of segment direction vector with coordinate axes
	if flops.Gt(math.Abs(m[1]*d[2]-m[2]*d[1]), b.HalfExtend[1]*adz+b.HalfExtend[2]*ady) ||
		flops.Gt(math.Abs(m[2]*d[0]-m[0]*d[2]), b.HalfExtend[0]*adz+b.HalfExtend[2]*adx) ||
		flops.Gt(math.Abs(m[0]*d[1]-m[1]*d[0]), b.HalfExtend[0]*ady+b.HalfExtend[1]*adx) {
		return false
	}
	// No separating axis found; segment must be overlapping AABB
	return true
}

// IntersectSegmentTriangle is given line pq and ccw triangle abc and return
// whether line pierces triangle. If so, also return the barycentric coordinates
// (u,v,w) of the intersection point.
func IntersectSegmentTriangle(p, q, a, b, c *glm.Vec3) (u, v, w float32, overlap bool) {
	pq := q.Sub(p)
	pa := a.Sub(p)
	pb := b.Sub(p)
	pc := c.Sub(p)
	// Test if pq is inside the edges bc, ca and ab. Done by testing
	// that the signed tetrahedral volumes, computed using scalar triple // products, are all positive
	u = glm.ScalarTripleProduct(&pq, &pc, &pb)
	if flops.Ltz(u) {
		return // already false and the other values are junk anyway.
	}
	v = glm.ScalarTripleProduct(&pq, &pa, &pc)
	if flops.Ltz(v) {
		return // already false and the other values are junk anyway.
	}
	w = glm.ScalarTripleProduct(&pq, &pb, &pa)
	if flops.Ltz(w) {
		return // already false and the other values are junk anyway.
	}
	// Compute the barycentric coordinates (u, v, w) determining the // intersection point r, r = u*a + v*b + w*c
	denom := 1.0 / (u + v + w)
	u *= denom
	v *= denom
	w *= denom // w = 1.0f - u - v;
	overlap = true
	return
}

// IntersectSegmentQuad is given line pq and ccw quadrilateral abcd, return
// whether the line pierces the triangle. If so, also return the point r of
// intersection.
func IntersectSegmentQuad(p, q, a, b, c, d *glm.Vec3) (glm.Vec3, bool) {
	pq := q.Sub(p)
	pa := a.Sub(p)
	pb := b.Sub(p)
	pc := c.Sub(p)
	// Determine which triangle to test against by testing against diagonal first
	m := pc.Cross(&pq)
	v := pa.Dot(&m) // glm.ScalarTripleProduct(pq, pa, pc);
	if flops.Gez(v) {
		// Test intersection against triangle abc
		u := -pb.Dot(&m) // glm.ScalarTripleProduct(pq, pc, pb);
		if flops.Ltz(u) {
			return glm.Vec3{}, false
		}
		w := glm.ScalarTripleProduct(&pq, &pb, &pa)
		if flops.Ltz(w) {
			return glm.Vec3{}, false
		}
		// Compute r, r = u*a + v*b + w*c, from barycentric coordinates (u, v, w)
		denom := 1.0 / (u + v + w)
		u *= denom
		v *= denom
		w *= denom // w = 1.0f - u - v;
		r := a.Mul(u)
		r.AddScaledVec(v, b)
		r.AddScaledVec(w, c)
		return r, true
	}
	// Test intersection against triangle dac
	pd := d.Sub(p)
	u := pd.Dot(&m) // glm.ScalarTripleProduct(pq, pd, pc);
	if u < 0 {
		return glm.Vec3{}, false
	}
	w := glm.ScalarTripleProduct(&pq, &pa, &pd)
	if w < 0 {
		return glm.Vec3{}, false
	}
	v = -v
	// Compute r, r = u*a + v*d + w*c, from barycentric coordinates (u, v, w)
	denom := 1.0 / (u + v + w)
	u *= denom
	v *= denom
	w *= denom // w = 1.0f - u - v;
	r := a.Mul(u)
	r.AddScaledVec(v, d)
	r.AddScaledVec(w, c)
	return r, true
}

// IntersectSegmentTriangle2 is given segment pq and triangle abc and returns
// whether segment intersects triangle and if so, also returns the barycentric
// coordinates (u,v,w) of the intersection point.
func IntersectSegmentTriangle2(p, q, a, b, c *glm.Vec3) (u, v, w, t float32, overlap bool) {
	ab := b.Sub(a)
	ac := c.Sub(a)
	qp := p.Sub(q)
	// Compute triangle normal. Can be precalculated or cached if
	// intersecting multiple segments against the same triangle
	n := ab.Cross(&ac)
	// Compute denominator d. If d <= 0, segment is parallel to or points
	// away from triangle, so exit early
	d := qp.Dot(&n)
	if d <= 0 {
		return
	}
	// Compute intersection t value of pq with plane of triangle. A ray
	// intersects iff 0 <= t. Segment intersects iff 0 <= t <= 1. Delay
	// dividing by d until intersection has been found to pierce triangle
	ap := p.Sub(a)
	t = ap.Dot(&n)
	if t < 0.0 {
		return
	}
	if t > d { // For segment; exclude this code line for a ray test
		return
	}
	// Compute barycentric coordinate components and test if within bounds
	e := qp.Cross(&ap)
	v = ac.Dot(&e)
	if v < 0.0 || v > d {
		return
	}
	w = -ab.Dot(&e)
	if w < 0.0 || v+w > d {
		return
	}
	// Segment/ray intersects triangle. Perform delayed division and // compute the last barycentric coordinate component
	ood := 1.0 / d
	t *= ood
	v *= ood
	w *= ood
	u = 1.0 - v - w
	overlap = true
	return
}

// IntersectSegmentCylinder intersects segment S(t)=sa+t(sb-sa), 0<=t<=1 against
// cylinder specified by p, q and r
func IntersectSegmentCylinder(sa, sb, p, q *glm.Vec3, r float32) (float32, bool) {
	const Epsilon = 0.00001
	d := q.Sub(p)
	m := sa.Sub(p)
	n := sb.Sub(sa)
	md := m.Dot(&d)
	nd := n.Dot(&d)
	dd := d.Dot(&d)
	// Test if segment fully outside either endcap of cylinder
	if md < 0.0 && md+nd < 0.0 {
		return 0, false
	} // Segment outside ’p’ side of cylinder
	if md > dd && md+nd > dd {
		return 0, false
	}
	// Segment outside ’q’ side of cylinder
	nn := n.Dot(&n)
	mn := m.Dot(&n)
	a := dd*nn - nd*nd
	k := m.Dot(&m) - r*r
	c := dd*k - md*md
	if math.Abs(a) < Epsilon {
		// Segment runs parallel to cylinder axis
		if c > 0 {
			return 0, false
		}
		// ’a’ and thus the segment lie outside cylinder
		// Now known that segment intersects cylinder; figure out how it intersects
		var t float32
		if md < 0 {
			t = -mn / nn // Intersect segment against ’p’ endcap
		} else if md > dd {
			t = (nd - mn) / nn // Intersect segment against ’q’ endcap
		} else {
			t = 0 // ’a’ lies inside cylinder
		}
		return t, true
	}
	b := dd*mn - nd*md
	discr := b*b - a*c
	if discr < 0 {
		return 0, false // No real roots; no intersection
	}
	t := (-b - math.Sqrt(discr)) / a
	if t < 0 || t > 1 {
		return 0, false // Intersection lies outside segment
	}
	if md+t*nd < 0 {
		// Intersection outside cylinder on ’p’ side
		if nd <= 0 {
			return 0, false // Segment pointing away from endcap
		}
		t = -md / nd
		// Keep intersection if Dot(S(t) - p, S(t) - p) <= r∧2
		return t, k+2*t*(mn+t*nn) <= 0
	} else if md+t*nd > dd {
		// Intersection outside cylinder on ’q’ side
		if nd >= 0 {
			return 0, false // Segment pointing away from endcap
		}
		t = (dd - md) / nd
		// Keep intersection if Dot(S(t) - q, S(t) - q) <= r∧2
		return t, k+dd-2*md+t*(2*(mn-nd)+t*nn) <= 0
	}
	// Segment intersects cylinder between the endcaps; t is correct
	return t, true
}
