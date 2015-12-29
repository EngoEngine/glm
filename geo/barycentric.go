package geo

import (
	"github.com/luxengine/glm"
)

// Barycentric returns the barycentric coordinates for a point in a triangle.
func Barycentric(a, b, c, p *glm.Vec3) (u, v, w float32) {
	v0, v1, v2 := b.Sub(a), c.Sub(a), p.Sub(a)
	d00 := v1.Dot(&v1)
	d01 := v0.Dot(&v1)
	d11 := v1.Dot(&v1)
	d20 := v2.Dot(&v0)
	d21 := v2.Dot(&v1)
	denom := 1 / (d00*d11 - d01*d01)
	v = (d11*d20 - d01*d21) * denom
	w = (d00*d21 - d01*d20) * denom
	u = 1 - v - w
	return
}
