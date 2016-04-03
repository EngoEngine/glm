package geo

import (
	"github.com/luxengine/glm"
	"github.com/luxengine/math"
)

// Sphere3 is a bounding volume for spheres in 3D.
type Sphere3 struct {
	Center          glm.Vec3
	Radius, Radius2 float32
}

// TestSphere3Sphere3 return true if the spheres overlap.
func TestSphere3Sphere3(a, b *Sphere3) bool {
	d := b.Center.Sub(&a.Center)
	l2 := d.Len2()
	r := a.Radius + b.Radius
	return l2 <= r*r
}

// AABB3FromSphere3 returns the AABB bounding this sphere.
//
// NOTE: If you need to use this function you better start questioning the
// algorithm you're implementing as the sphere is both faster and bounds the
// underlying object better.
func AABB3FromSphere3(s *Sphere3) AABB3 {
	return AABB3{
		Center: s.Center,
		Radius: glm.Vec3{s.Radius, s.Radius, s.Radius},
	}
}

// MergeSphere3Point updates the bounding sphere to encompass v if needed.
func MergeSphere3Point(s *Sphere3, v *glm.Vec3) {
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

// EigenSphere3 sets this sphere to the bounding sphere of the given points using
// eigen values algorithm, this doesn't necessarily wrap all the points so use
// RitterEigenSphere.
func EigenSphere3(points []glm.Vec3) Sphere3 {
	var s Sphere3
	var m, v glm.Mat3

	// Compute the covariance matrix m
	CovarianceMatrix3(&m, points)

	// Decompose it into eigenvectors (in v) and eigenvalues (in m)
	Jacobi(&m, &v)

	// Find the component with largest magnitude eigenvalue (largest spread)
	var e glm.Vec3

	var maxc int

	var maxf float32
	maxe := math.Abs(m[0])

	if maxf = math.Abs(m[3*1+1]); maxf > maxe {
		maxc = 1
		maxe = maxf
	}
	if maxf = math.Abs(m[3*2+2]); maxf > maxe {
		maxc = 2
		maxe = maxf
	}
	e[0] = v[3*maxc+0]
	e[1] = v[3*maxc+1]
	e[2] = v[3*maxc+2]

	// Find the most extreme points along direction ’e’
	imin, imax := ExtremePointsAlongDirection3(&e, points)
	minpt := points[imin]
	maxpt := points[imax]
	u := maxpt.Sub(&minpt)
	dist := u.Len()
	s.Radius = dist * 0.5

	t := minpt.Add(&maxpt)
	s.Center.MulOf(0.5, &t)
	return s
}

// RitterEigenSphere sets this sphere to wrap all the given points using eigen
// values algorithm.
func RitterEigenSphere(points []glm.Vec3) Sphere3 {
	// Start with sphere from maximum spread
	s := EigenSphere3(points)
	// Grow sphere to include all points
	for i := range points {
		MergeSphere3Point(&s, &points[i])
	}
	return s
}
