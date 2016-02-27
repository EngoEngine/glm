package geo

import (
	"github.com/luxengine/glm"
	"github.com/luxengine/math"
)

// Barycentric returns the barycentric coordinates for a point in a triangle.
// The barycentric coordinates cna be used for interpolation of {a, b, c} to
// point p (such as normals, texture coordinates, colors, etc)
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

// BarycentricCache is a cache of data for quick Barycentric calls. Use it if
// you need to call Barycentric on the same triangle over and over.
type BarycentricCache struct {
	a, v0, v1            glm.Vec3
	d00, d01, d11, denom float32
}

// BarycentricCacheFromTriangle takes a triangle and returns a  barycentric
// cache for queries.
func BarycentricCacheFromTriangle(a, b, c *glm.Vec3) BarycentricCache {
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

// Barycentric returns the barycentric coordinates for a point in a triangle.
func (b *BarycentricCache) Barycentric(p *glm.Vec3) (u, v, w float32) {
	v2 := p.Sub(&b.a)

	d20 := v2.Dot(&b.v0)
	d21 := v2.Dot(&b.v1)

	v = (b.d11*d20 - b.d01*d21) * b.denom
	w = (b.d00*d21 - b.d01*d20) * b.denom
	u = 1 - v - w

	return
}

// IsPointInTriangle returns true if the point p projected on triangle {a,b,c}
// is inside triangle {a,b,c}
func IsPointInTriangle(p, a, b, c *glm.Vec3) bool {
	_, v, w := Barycentric(a, b, c, p)
	return v >= 0 && w >= 0 && (v+w) <= 1
}

func triArea2D(x1, y1, x2, y2, x3, y3 float32) float32 {
	return (x1-x2)*(y2-y3) - (x2-x3)*(y1-y2)
}

// candidate for optimized barycentric. Actually fails so far. But weird
// inlining might be at fault
func barycentric2(a, b, c, p *glm.Vec3) (u, v, w float32) {
	// Unnormalized triangle normal
	//bma := glm.Vec3{b[0] - a[0], b[1] - a[1], b[2] - a[2]}
	//cma := glm.Vec3{c[0] - a[0], c[1] - a[1], c[2] - a[2]}
	bma, cma := b.Sub(a), c.Sub(a)
	m := bma.Cross(&cma)

	// Nominators and one-over-denominator for u and v ratios
	var nu, nv, ood float32

	// Absolute components for determining projection plane
	x, y, z := math.Abs(m[0]), math.Abs(m[1]), math.Abs(m[2])

	// Compute areas in plane of largest projection
	if x >= y && x >= z {
		// x is largest, project to the yz plane
		nu = triArea2D(p[1], p[2], b[1], b[2], c[1], c[2]) // Area of PBC in yz plane
		nv = triArea2D(p[1], p[2], c[1], c[2], a[1], a[2]) // Area of PCA in yz plane
		ood = 1 / m[0]                                     // 1/(2*area of ABC in yz plane)
	} else if y >= x && y >= z {
		// y is largest, project to the xz plane
		nu = triArea2D(p[0], p[2], b[0], b[2], c[0], c[2])
		nv = triArea2D(p[0], p[2], c[0], c[2], a[0], a[2])
		ood = 1 / -m[1]
	} else {
		// z is largest, project to the xy plane
		nu = triArea2D(p[0], p[1], b[0], b[1], c[0], c[1])
		nv = triArea2D(p[0], p[1], c[0], c[1], a[0], a[1])
		ood = 1 / m[2]
	}

	u = nu * ood
	v = nv * ood
	w = 1 - u - v
	return
}
