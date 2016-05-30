package geo

import (
	"fmt"
	"github.com/engoengine/glm"
	"strconv"
)

// Simplex represents a simplex in 3D. Either a point, a line, a triangle,
// or a tetrahedron.
type Simplex struct {
	// the Points contained in the simplex. Data past Points[Size] is assumed to
	// be garbage.
	Points [4]glm.Vec3 // use an array to keep the memory all in 1 spot
	// the current extend of the simplex.
	Size int
}

// Merge merges the given vector to the simplex. This will panic if you add a
// 5th vertex.
func (s *Simplex) Merge(u *glm.Vec3) {
	s.Points[s.Size] = *u
	s.Size++
}

// NearestToOrigin modifies the simplex to contain only the minimum amount of
// points required to describe the direction to origin, it also returns the next
// direction to search in GJK and true if the origin is contained in the simplex
func (s *Simplex) NearestToOrigin() (direction glm.Vec3, containsOrigin bool) {
	// TODO(hydroflame): implement
	const (
		a = iota
		b
		c
		d
	)

	switch s.Size {
	case 4:
		ab := s.Points[b].Sub(&s.Points[a])
		ac := s.Points[c].Sub(&s.Points[a])
		ad := s.Points[d].Sub(&s.Points[a])
		ap := s.Points[a].Inverse()

		// Check if Origin is in vertex region outside A
		d1, d2, d3 := ab.Dot(&ap), ac.Dot(&ap), ad.Dot(&ap)
		if d1 <= 0 && d2 <= 0 && d3 <= 0 {
			// inside voronoi region of A
			fmt.Print("		A")
			return glm.Vec3{}, false
		}

		bp := s.Points[b].Inverse()
		d4, d5, d6 := ab.Dot(&bp), ac.Dot(&bp), ad.Dot(&bp)
		if d4 >= 0 && d5 <= d4 && d6 <= d4 {
			// inside voronoi region of B
			fmt.Print("		B")
			return glm.Vec3{}, false
		}

		cp := s.Points[c].Inverse()
		d7, d8, d9 := ac.Dot(&cp), ab.Dot(&cp), ad.Dot(&cp)
		if d7 >= 0 && d8 <= d7 && d9 <= d7 {
			// inside voronoi region of C
			fmt.Print("		C")
			return glm.Vec3{}, false
		}

		dp := s.Points[d].Inverse()
		d10, d11, d12 := ad.Dot(&dp), ab.Dot(&dp), ac.Dot(&dp)
		if d10 >= 0 && d11 <= d10 && d12 <= d10 {
			// inside voronoi region of D
			fmt.Print("		D")
			return glm.Vec3{}, false
		}

		vab := d1*d5 - d2*d6 - d3*d4
		if vab <= 0 && d1 >= 0 && d4 <= 0 {
			fmt.Print("		yes")
			fmt.Print("	", s.Points)
		}
	case 3:
	case 2:
	case 1:
	default: //case Size < 1 || Size > 4
		panic("Simplex.Size=" + strconv.Itoa(int(s.Size)) + ", need 1, 2, 3, or 4")
	}

	return glm.Vec3{}, false
}
