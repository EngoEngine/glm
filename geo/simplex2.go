package geo

import (
	"github.com/luxengine/glm"
)

// Simplex2 represents a simplex in 2D. Either a point, a line, or a triangle.
// We use this in GJK to explore the Minskowski difference.
type Simplex2 struct {
	Points [3]glm.Vec2 // use an array to keep the memory al in 1 spot
	Size   int
}

// Simplex2MergeVec merges the given vector to the simplex. This will panic if you add a
// 4th vertex.
func Simplex2MergeVec(s *Simplex2, u *glm.Vec2) {
	s.Points[s.Size] = *u
	s.Size++
}

// Simplex2NearestToOrigin modifies the simplex to contain only the minimum amount of
// points required to describe the direction to origin, it also returns the next
// direction to search in GJK and true if the origin is contained in the simplex
func Simplex2NearestToOrigin(s *Simplex2) (direction glm.Vec2, containsOrigin bool) {
	if s.Size == 3 {
		ab := s.Points[1].Sub(&s.Points[0])
		ac := s.Points[2].Sub(&s.Points[0])
		ap := s.Points[0].Inverse()

		// Check if Origin is in vertex region outside A
		d1, d2 := ab.Dot(&ap), ac.Dot(&ap)
		if d1 <= 0 && d2 <= 0 {
			var zero glm.Vec2
			s.Size = 1
			return ap, s.Points[0].ApproxEqual(&zero)
		}

		// Check if Origin is in vertex region outside B
		bp := s.Points[1].Inverse()
		d3, d4 := ab.Dot(&bp), ac.Dot(&bp)
		if d3 >= 0 && d4 <= d3 {
			var zero glm.Vec2
			s.Size = 1
			s.Points[0] = s.Points[1]
			return bp, s.Points[0].ApproxEqual(&zero)
		}

		// Check if Origin is in edge region of AB, if so return projection of
		// Origin onto AB
		vc := d1*d4 - d2*d3
		if vc <= 0 && d1 >= 0 && d3 <= 0 {
			s.Size = 2
			ret := glm.Vec2{-ab[1], ab[0]}
			if ret.Dot(&ac) < 0 {
				return ret, glm.FloatEqual(ab.Cross(&ap), 0)
			}
			return ret.Inverse(), glm.FloatEqual(ab.Cross(&ap), 0)
		}

		// Check if P in vertex region outside C
		cp := s.Points[2].Inverse()
		d5, d6 := ab.Dot(&cp), ac.Dot(&cp)
		if d6 >= 0 && d5 <= d6 {
			var zero glm.Vec2
			s.Size = 1
			s.Points[0] = s.Points[2]
			return cp, s.Points[0].ApproxEqual(&zero)
		}

		// Check if Origin is in edge region of AC, if so return projection of
		// Origin onto AC
		vb := d5*d2 - d1*d6
		if vb <= 0 && d2 >= 0 && d6 <= 0 {
			s.Size = 2
			s.Points[1] = s.Points[2]
			ret := glm.Vec2{ac[1], -ac[0]}
			if ret.Dot(&ab) < 0 {
				return ret, glm.FloatEqual(ac.Cross(&ap), 0)
			}
			return ret.Inverse(), glm.FloatEqual(ac.Cross(&ap), 0)
		}

		// Check if Origin is in edge region of BC, if so return projection of
		// Origin onto BC
		va := d3*d6 - d5*d4
		if va <= 0 && (d4-d3) >= 0 && (d5-d6) >= 0 {
			bc := s.Points[2].Sub(&s.Points[1])
			s.Size = 2
			s.Points[0] = s.Points[1]
			s.Points[1] = s.Points[2]
			ib := s.Points[1].Inverse()

			ret := glm.Vec2{-bc[1], bc[0]}
			if ret.Dot(&ab) > 0 {
				return ret, glm.FloatEqual(bc.Cross(&ib), 0)
			}
			return ret.Inverse(), glm.FloatEqual(bc.Cross(&ib), 0)
		}

		return glm.Vec2{}, true
	}

	if s.Size == 2 {
		zto := s.Points[1].Sub(&s.Points[0]) // zero to one
		i0, i1 := s.Points[0].Inverse(), s.Points[1].Inverse()

		if zto.Dot(&i1) > 0 { // in voronoi zone of Points[1]
			var zero glm.Vec2
			s.Points[0] = s.Points[1]
			s.Size = 1
			return i1, s.Points[0].ApproxEqual(&zero)
		}

		if zto.Dot(&i0) < 0 { // in voronoi zone of Points[0]
			var zero glm.Vec2
			s.Size = 1
			return i0, s.Points[0].ApproxEqual(&zero)
		}

		// check if the origin is on the line
		perp := glm.Vec2{-i0[1], i0[0]}
		if glm.FloatEqual(perp.Dot(&i1), 0) {
			return perp, true
		}

		perp = glm.Vec2{-zto[1], zto[0]}
		if perp.Dot(&i0) > 0 && perp.Dot(&i1) > 0 {
			return perp, false
		}
		return perp.Inverse(), false
	}

	var zero glm.Vec2
	if s.Points[0].ApproxEqual(&zero) {
		return glm.Vec2{}, true
	}

	return s.Points[0].Inverse(), false
}
