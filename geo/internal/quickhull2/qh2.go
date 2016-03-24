package quickhull2

import (
	"github.com/luxengine/glm"
	"github.com/luxengine/math"
)

// SupportDirection is the direction that we use to cache the direction for the
// Support function.
var SupportDirection = [3]glm.Vec2{
	glm.Vec2{math.Cos(0 * math.Pi / 3), math.Sin(0 * math.Pi / 3)},
	glm.Vec2{math.Cos(2 * math.Pi / 3), math.Sin(2 * math.Pi / 3)},
	glm.Vec2{math.Cos(4 * math.Pi / 3), math.Sin(4 * math.Pi / 3)},
}

// Vertex is a convex hull vertex
type Vertex struct {
	Position   glm.Vec2
	Prev, Next int // index in the vertex slice
}

// Edge is a convex hull edge
type Edge struct {
	Vertices           [2]int
	PrevEdge, NextEdge int // index in the edge slice
}

/*pseudo GJK

  function GJK_intersection(shape p, shape q, vector initial_axis):
      vector  A = Support(p, initial_axis) - Support(q, -initial_axis)
      simplex s = {A}
      vector  D = -A
      loop:
          A = Support(p, D) - Support(q, -D)
          if dot(A, D) < 0:
             reject
          s = s ∪ A
          s, D, contains_origin = NearestSimplex(s)
          if contains_origin:
             accept

*/
