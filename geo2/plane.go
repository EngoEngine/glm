package geo2

import (
	"github.com/luxengine/glm"
)

// Plane is a hyperplane in 2D.
type Plane struct {
	// The normal to the plane.
	N glm.Vec2

	// The arbitrary point that the plane starts on.
	P glm.Vec2
}

// PlaneFromPoints computes the plane given by (a,b).
func PlaneFromPoints(a, b *glm.Vec2) Plane {
	w := b.Sub(a)
	w.Normalize()
	w[0], w[1] = -w[1], w[0]

	return Plane{
		N: w,
		P: *a,
	}
}

// DistanceToPlane returns the distance of v to plane p.
func DistanceToPlane(p *Plane, v *glm.Vec2) float32 {
	u := v.Sub(&p.P)
	return u.Dot(&p.N)
}
