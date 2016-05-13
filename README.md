# glm [![Build Status](http://lionheart.hydroflame.net:8080/job/glm/lastSuccessfulBuild/artifact/status.svg)](http://lionheart.hydroflame.net:8080/job/glm/) [![Tests](http://lionheart.hydroflame.net:8080/job/glm/lastSuccessfulBuild/artifact/test.svg)](http://lionheart.hydroflame.net:8080/job/glm/) [![Coverage](http://lionheart.hydroflame.net:8080/job/glm/lastSuccessfulBuild/artifact/cover.svg)](http://lionheart.hydroflame.net:8080/job/glm/) [![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/luxengine/glm)

VENDOR THIS IF YOU'RE USING IT. The API is not quite stable yet. Specialy geo*.

More efficient version then go-gl math lib and better name (mgl32 is too long to type).

The problem with go-gl implementation is that every operation returns a new matrix/quaternion/vector, you can't reuse memory. Benchmark reflect that this causes quite the slowdown. see [issue 29](https://github.com/go-gl/mathgl/issues/29). This library is a fork of mgl32 but a ton of methods we're added in order to allow the user to have more control over the memory. First, most methods take pointer argument and second there is more then 1 method to do the same operation.

This library uses lux math (native float32 math) instead of the standard library math. 

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

