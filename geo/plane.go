package geo

import (
	"github.com/luxengine/glm"
)

// Plane is a hyperplane in 3D.
type Plane struct {
	// the normal to the plane
	N glm.Vec3

	// The arbitrary point that the plane starts on.
	P glm.Vec3
}

// PlaneFromPoints computes the plane given by (a,b,c), ordered ccw.
func PlaneFromPoints(a, b, c *glm.Vec3) Plane {
	v0, v1 := b.Sub(a), c.Sub(a)
	n := v0.Cross(&v1)
	n.Normalize()
	return Plane{
		N: n,
		P: *a,
	}
}

// DistanceToPlane returns the distance of v to plane p.
func DistanceToPlane(p *Plane, v *glm.Vec3) float32 {
	u := v.Sub(&p.P)
	return u.Dot(&p.N)
}
