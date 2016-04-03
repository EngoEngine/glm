package geo

import (
	"github.com/luxengine/glm"
)

// Plane3 is a hyperplane in 3D.
type Plane3 struct {
	N glm.Vec3
	D float32
}

// Plane3FromPoints computes the plane given by (a,b,c), ordered ccw.
func Plane3FromPoints(a, b, c *glm.Vec3) Plane3 {
	v0, v1 := b.Sub(a), c.Sub(a)
	n := v0.Cross(&v1)
	n.Normalize()
	return Plane3{
		N: n,
		D: n.Dot(a),
	}
}

// DistanceToPlane3 returns the distance of v to plane p.
func DistanceToPlane3(p *Plane3, v *glm.Vec3) float32 {
	// transform the point and the plane such that the plane pass by {0,0,0}
	w := glm.Vec3{v[0], v[1] - p.D, v[2]}
	return p.N.Dot(&w)
}

// Plane2 is a hyperplane in 2D.
type Plane2 struct {
	N glm.Vec2
	D float32
}

// Plane2FromPoints computes the plane given by (a,b).
func Plane2FromPoints(a, b *glm.Vec2) Plane2 {
	w := b.Sub(a)
	w.Normalize()
	w[0], w[1] = -w[1], w[0]

	return Plane2{
		N: w,
		D: w.Dot(a),
	}
}

// DistanceToPlane2 returns the distance of v to plane p.
func DistanceToPlane2(p *Plane2, v *glm.Vec2) float32 {
	w := glm.Vec2{v[0], v[1] - p.D}
	return p.N.Dot(&w)
}
