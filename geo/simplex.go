package geo

import (
	"github.com/luxengine/glm"
)

// Simplex2 represents a simple in 2d (so either a point, a line, or a triangle)
type Simplex2 struct {
	points [3]glm.Vec3 // use an array to keep the memory al in 1 spot
	size   int
}

// NearestToOrigin modifies the simplex to contain only the minimum amount of
// points required to describe the direction to origin, returns the next
// direction to search in GJK and true if the origin is contained in the simplex
func (s *Simplex2) NearestToOrigin() (direction glm.Vec2, containsOrigin bool) {
	if s.size == 3 {
		ab, ac, ap := s.points[1].Sub(&s.points[0]), s.points[2].Sub(&s.points[0]), s.points[0].Inverse()

		// Check if Origin is in vertex region outside A
		d1, d2 := ab.Dot(&ap), ac.Dot(&ap)
		if d1 <= 0 && d2 <= 0 {
			var zero glm.Vec2
			s.size = 1
			return ap, s.points[0].ApproxEqual(&zero)
		}

		// Check if Origin is in vertex region outside B
		bp := s.points[1].Inverse()
		d3, d4 := ab.Dot(&bp), ac.Dot(&ap)
		if d3 >= 0 && d4 <= d3 {
			var zero glm.Vec2
			s.size = 1
			s.points[0] = s.points[1]
			return bp, s.points[0].ApproxEqual(&zero)
		}

		// Check if Origin is in edge region of AB, if so return projection of
		// Origin onto AB
		vc := d1*d4 - d3*d2
		if vc <= 0 && d1 >= 0 && d3 <= 0 {
			s.size = 2
			return glm.Vec2{-ab[1], ab[0]}, false // TODO check if return false is correct
		}

		// Check if P in vertex region outside C
		cp := s.points[2].Inverse()
		d5, d6 := ab.Dot(&cp), ac.Dot(&cp)
		if d6 >= 0 && d5 <= d6 {
			var zero glm.Vec2
			s.size = 1
			s.points[0] = s.points[2]
			return cp, s.points[0].ApproxEqual(&zero)
		}

		// Check if Origin is in edge region of AC, if so return projection of
		// Origin onto AC
		vb := d5*d2 - d1*d6
		if vb <= 0 && d2 >= 0 && d6 <= 0 {
			s.size = 2
			s.points[1] = s.points[2]
			return glm.Vec2{-ac[1], ac[0]}, false // TODO check if return false is correct
		}

		// Check if Origin is in edge region of BC, if so return projection of
		// Origin onto BC
		va := d3*d6 - d5*d4
		if va <= 0 && (d4-d3) >= 0 && (d5-d6) >= 0 {
			s.size = 2
			s.points[0] = s.points[1]
			s.points[1] = s.points[2]
			return glm.Vec2{-bc[1], bc[0]}, false // TODO check if return false is correct
		}

		return glm.Vec2, true
	}

	if s.size == 2 {
		zto := s.points[1].Sub(&s.points[0]) // zero to one
		i0, i1 := s.points[0].Inverse(), s.points[1].Inverse()

		if zto.Dot(&i1) > 0 { // in voronoi zone of points[1]
			var zero glm.Vec2
			s.points[0] = s.points[1]
			s.size = 1
			return i1, s.points[0].ApproxEqual(&zero)
		}

		if zto.Dot(&i0) < 0 { // in voronoi zone of points[0]
			var zero glm.Vec2
			s.size = 1
			return i0, s.points[0].ApproxEqual(&zero)
		}

		// check if the origin is on the line
		perp := glm.Vec2{-i0[1], i0[0]}
		if glm.FloatEqual(perp.Dot(&i1), 0) {
			return perp, true
		}

		perp = glm.Vec2{-zto[1], zto[0]}
		if perp.Dot(&i0) > 0 && perp.Dot(&i1) {
			return perp, false
		}
		return perp.Inverse(), false
	}

	if s.points[0].ApproxEqual(glm.Vec2{}) {
		return glm.Vec2{}, true
	}

	return s.points[0].Inverse(), false
}
