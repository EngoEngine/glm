// Package glm declares many linear algrebric types to do computer geometry. The
// matrices are column major (same as OpenGL) and all the angles are in radians.
// It is based off github.com/go-gl/mathgl version but has a very different API
// to take advantage of cache locality. This makes this library generally faster
// to execute but harder to read and write. It also uses a native math library
// for float32.
package glm
