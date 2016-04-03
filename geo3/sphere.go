package geo3

import (
	"github.com/luxengine/glm"
	"github.com/luxengine/math"
)

// Sphere is a bounding volume for spheres.
type Sphere struct {
	Center          glm.Vec3
	Radius, Radius2 float32
}

// TestSphereSphere return true if the spheres overlap.
func TestSphereSphere(a, b *Sphere) bool {
	d := b.Center.Sub(&a.Center)
	l2 := d.Len2()
	r := a.Radius + b.Radius
	return l2 <= r*r
}

// AABBFromSphere returns the AABB bounding this sphere.
//
// NOTE: If you need to use this function you better start questioning the
// algorithm you're implementing as the sphere is both faster and bounds the
// underlying object better.
func AABBFromSphere(s *Sphere) AABB {
	return AABB{
		Center: s.Center,
		Radius: glm.Vec3{s.Radius, s.Radius, s.Radius},
	}
}

// MergePoint updates the bounding sphere to encompass v if needed.
func (s *Sphere) MergePoint(v *glm.Vec3) {
	// Compute squared distance between point and sphere center
	d := v.Sub(&s.Center)
	dist2 := d.Len2()
	// Only update s if point p is outside it
	if dist2 > s.Radius*s.Radius {
		dist := math.Sqrt(dist2)
		newRadius := (s.Radius + dist) * 0.5
		k := (newRadius - s.Radius) / dist
		s.Radius = newRadius
		s.Center.AddScaledVec(k, &d)
		s.Radius2 = s.Radius * s.Radius
	}
}

// EigenSphere sets this sphere to the bounding sphere of the given points using
// eigen values algorithm, this doesn't necessarily wrap all the points so use
// RitterEigenSphere.
func EigenSphere(points []glm.Vec3) Sphere {
	var m glm.Mat3

	// Compute the covariance matrix m
	CovarianceMatrix(&m, points)

	var v glm.Mat3
	// Decompose it into eigenvectors (in v) and eigenvalues (in m)
	Jacobi(&m, &v)

	// Find the component with largest magnitude eigenvalue (largest spread)
	maxe := math.Abs(m[0])

	var maxc int
	if maxf := math.Abs(m[3*1+1]); maxf > maxe {
		maxc = 1
		maxe = maxf
	}
	if maxf := math.Abs(m[3*2+2]); maxf > maxe {
		maxc = 2
		maxe = maxf
	}

	var e glm.Vec3

	e[0] = v[3*maxc+0]
	e[1] = v[3*maxc+1]
	e[2] = v[3*maxc+2]

	// Find the most extreme points along direction e
	imin, imax := ExtremePointsAlongDirection(&e, points)
	minpt := points[imin]
	maxpt := points[imax]
	u := maxpt.Sub(&minpt)
	dist := u.Len()

	var s Sphere
	s.Radius = dist * 0.5

	t := minpt.Add(&maxpt)
	s.Center.MulOf(0.5, &t)
	return s
}

// RitterEigenSphere sets this sphere to wrap all the given points using eigen
// values as base.
func RitterEigenSphere(points []glm.Vec3) Sphere {
	// Start with sphere from maximum spread
	s := EigenSphere(points)
	// Grow sphere to include all points
	for i := range points {
		s.MergePoint(&points[i])
	}
	return s
}
