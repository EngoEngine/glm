package geo2

import (
	"github.com/luxengine/glm"
)

// BarycentricCache is a cache of data for quick Barycentric calls. Use it if
// you need to call Barycentric on the same triangle over and over.
type BarycentricCache struct {
	a, v0, v1            glm.Vec2
	d00, d01, d11, denom float32
}

// Barycentric returns the barycentric coordinates for a point in a triangle.
// The barycentric coordinates cna be used for interpolation of {a, b, c} to
// point p (such as normals, texture coordinates, colors, etc).
func Barycentric(a, b, c, p *glm.Vec2) (u, v, w float32) {
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

// BarycentricWithCache returns the barycentric coordinates for a point in a
// triangle.
func BarycentricWithCache(b *BarycentricCache, p *glm.Vec2) (u, v, w float32) {
	v2 := p.Sub(&b.a)

	d20 := v2.Dot(&b.v0)
	d21 := v2.Dot(&b.v1)

	v = (b.d11*d20 - b.d01*d21) * b.denom
	w = (b.d00*d21 - b.d01*d20) * b.denom
	u = 1 - v - w

	return
}

// BarycentricCacheFromTriangle takes a triangle and returns a barycentric
// cache for queries.
func BarycentricCacheFromTriangle(a, b, c *glm.Vec2) BarycentricCache {
	v0, v1 := b.Sub(a), c.Sub(a)
	d00 := v1.Dot(&v1)
	d01 := v0.Dot(&v1)
	d11 := v1.Dot(&v1)

	return BarycentricCache{
		a:     *a,
		v0:    v0,
		v1:    v1,
		d00:   d00,
		d01:   d01,
		d11:   d11,
		denom: 1 / (d00*d11 - d01*d01),
	}
}
