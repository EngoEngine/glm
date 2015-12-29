package geo

import (
	"github.com/luxengine/glm"
)

// Plane represents a plane in 3d space
type Plane struct {
	N glm.Vec3
	D float32
}

// PlaneFromPoints computes the plane given by (a,b,c), ordered counter
// clockwise.
func PlaneFromPoints(a, b, c *glm.Vec3) Plane {
	v0, v1 := b.Sub(a), c.Sub(a)
	n := v0.Cross(&v1)
	n.Normalize()
	return Plane{
		N: n,
		D: n.Dot(a),
	}
}

// DistanceToPlane returns the distance of v to plane p.
func (p *Plane) DistanceToPlane(v *glm.Vec3) float32 {
	w := glm.Vec3{v.X(), v.Y() - p.D, v.Z()}
	return p.N.Dot(&w)
}
