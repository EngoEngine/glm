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

// Intersects return true if the spheres overlap.
func (a *Sphere3) Intersects(b *Sphere3) bool {
	d := b.Center.Sub(&a.Center)
	l2 := d.Len2()
	r := a.Radius + b.Radius
	return l2 <= r*r
}

// AABB3 returns the AABB bounding this sphere.
//
// NOTE: If you need to use this function you better start questioning the
// algorithm you're implementing as the sphere is both faster and bounds the
// underlying object better.
func (a *Sphere3) AABB3() AABB3 {
	return AABB3{
		Center: a.Center,
		Radius: glm.Vec3{a.Radius, a.Radius, a.Radius},
	}
}

// OfSphereAndPt updates the bounding sphere to encompass v if needed.
func (a *Sphere3) OfSphereAndPt(v *glm.Vec3) {
	// Compute squared distance between point and sphere center
	d := v.Sub(&a.Center)
	dist2 := d.Len2()
	// Only update s if point p is outside it
	if dist2 > a.Radius*a.Radius {
		dist := math.Sqrt(dist2)
		newRadius := (a.Radius + dist) * 0.5
		k := (newRadius - a.Radius) / dist
		a.Radius = newRadius
		a.Center.AddScaledVec(k, &d)
		a.Radius2 = a.Radius * a.Radius
	}
}

// EigenSphere sets this sphere to the bounding sphere of the given points using
// eigen values algorithm, this doesn't necessarily wrap all the points so use
// RitterEigenSphere.
func (a *Sphere3) EigenSphere(points []glm.Vec3) {
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
	s := maxpt.Sub(&minpt)
	dist := s.Len()
	a.Radius = dist * 0.5

	t := minpt.Add(&maxpt)
	a.Center.MulOf(0.5, &t)
}

// RitterEigenSphere sets this sphere to wrap all the given points using eigen
// values algorithm.
func (a *Sphere3) RitterEigenSphere(points []glm.Vec3) {
	// Start with sphere from maximum spread
	a.EigenSphere(points)
	// Grow sphere to include all points
	for i := range points {
		a.OfSphereAndPt(&points[i])
	}
}
