# glm ![Build Status](http://jenkins.hydroflame.net/job/glm/lastBuild/artifact/status.svg) ![Tests](http://jenkins.hydroflame.net/job/glm/lastBuild/artifact/tests.svg) ![Coverage](http://jenkins.hydroflame.net/job/glm/lastBuild/artifact/coverage.svg)
[![GoDoc](https://godoc.org/github.com/luxengine/glm?status.svg)](https://godoc.org/github.com/luxengine/glm)

More efficient version then go-gl math lib and better name (math32 is too long to type).

The problem with go-gl implementation is that every operation returns a new matrix, you can reuse memory. Benchmark reflect that this causes quite the slowdown. see [issue 29](https://github.com/go-gl/mathgl/issues/29) on mathgl github. This library is basically a copy of mgl32 but a ton of methods we're added in order to allow the user to have more control over the memory. First, most methods take pointer argument and second there is more then 1 method to do the same operation.

This library uses lux math, so essentially everytime a new pure float32 function comes out this library will get faster as well. 

In the future, when we have more knowledge of plan9 we intend to insert some SIMD operations for the more hardcore stuff.
```Go
func (m1 *Mat2) Add(m2 *Mat2) *Mat2 {
	return &Mat2{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3]}
}

func (m1 *Mat2) SumOf(m2, m3 *Mat2) *Mat2 {
	m1[0] = m2[0] + m3[0]
	m1[1] = m2[1] + m3[1]
	m1[2] = m2[2] + m3[2]
	m1[3] = m2[3] + m3[3]
	return m1
}

func (m1 *Mat2) SumWith(m2 *Mat2) *Mat2 {
	m1[0] += m2[0]
	m1[1] += m2[1]
	m1[2] += m2[2]
	m1[3] += m2[3]
	return m1
}
```

`nameofop` takes the 2 elements and does `op` with them, storing the result in a new element. `x1 := x2 op x3`  
`nameofopOf` takes 3 argument, does the operation on the last 2 and stores the result in the first `x1 = x2 op x3`  
`nameofopWith` takes 2 element, does `op` with them and stores the results in the first. `x1 = x1 op x2`

