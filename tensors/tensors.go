package tensors

import (
	"github.com/luxengine/glm"
)

// Continuous returns a continuous inertia tensor from the given coefficients.
//	[ Ix , -Ixy, -Ixz]
//	[-Ixy,  Iy , -Iyz]
//	[-Ixz, -Iyz,   Iz]
func Continuous(ix, iy, iz, ixy, ixz, iyz float32) glm.Vec3 {
	return glm.Vec3{
		ix - ixy, -ixz,
		-ixy, iy, -iyz,
		-ixz, -iyz, iz,
	}
}

// Sphere returns the inertia tensor of a sphere.
func Sphere(mass, radius float32) glm.Mat3 {
	v := 0.4 * mass * radius * radius
	return glm.Mat3{
		v, 0, 0,
		0, v, 0,
		0, 0, v,
	}
}

// Cuboid returns the inertiatensor of a cuboid shape. dx, dy, dz
// are all full sizes as if the box was axis aligned.
func Cuboid(mass, dx, dy, dz float32) glm.Mat3 {
	const (
		v = 1.0 / 12.0
	)
	dx2, dy2, dz2 := dx*dx, dy*dy, dz*dz
	return glm.Mat3{
		f * mass * (dy2 + dz2), 0, 0,
		0, f * mass * (dx2 + dz2), 0,
		0, 0, f * mass * (dx2 + dy2),
	}
}

// Cylinder returns the inertia tensor of a cylinder whose principal axe is
// along the Z axis.
func Cylinder(mass, radius, height float32) glm.Mat3 {
	qmr2 := mass * radius * radius / 4
	tmh2 := mass * height * height / 12
	return glm.Mat3{
		tmh2 + qmr2, 0, 0,
		0, tmh2 + qmr2, 0,
		0, 0, qmr2 * 2,
	}
}

// Cone returns the inertia tensor of a cone whose principal axe is along the Z
// axis.
func Cone(mass, radius, height float32) glm.Mat3 {
	toemh2 := 3 * mass * height * height / 80
	totmr2 := 3 * mass * radius * radius / 20
	return glm.Mat3{
		toemh2 + totmr2, 0, 0,
		0, toemh2 + totmr2, 0,
		0, 0, totmr2 * 2,
	}
}
