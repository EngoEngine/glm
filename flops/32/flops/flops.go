package flops

import (
	"github.com/engoengine/math"
)

const (
	epsilon = 0.000001
)

// refequal returns true if the floats are approximatelly equal. this function
// is used as reference for the unwrapped equal.
func refequal(a, b float32) bool {
	return math.Abs(a-b) <= epsilon*math.Max(math.Max(1, math.Abs(a)), math.Abs(b))
}

// Eq returns true if the floats are approximatelly equal.
func Eq(a, b float32) bool {
	if a > 0 {
		if b > 0 {
			if a > b {
				if a > 1 {
					return a-b <= epsilon*a
				}
				return a-b <= epsilon
			}
			if b > 1 {
				return b-a <= epsilon*b
			}
			return b-a <= epsilon
		}
		return false
	}
	if b > 0 {
		return false
	}
	if a > b {
		if b < -1 {
			return -(b - a) <= epsilon*-b
		}
		return -(b - a) <= epsilon
	}
	if a < -1 {
		return -(a - b) <= epsilon*-a
	}
	return -(a - b) <= epsilon
}

// Ne returns true if the floats are not approximately equal
func Ne(a, b float32) bool {
	if a > 0 {
		if b > 0 {
			if a > b {
				if a > 1 {
					return a-b > epsilon*a
				}
				return a-b > epsilon
			}
			if b > 1 {
				return b-a > epsilon*b
			}
			return b-a > epsilon
		}
		return true
	}
	if b > 0 {
		return true
	}
	if a > b {
		if b < -1 {
			return -(b - a) > epsilon*-b
		}
		return -(b - a) > epsilon
	}
	if a < -1 {
		return -(a - b) > epsilon*-a
	}
	return -(a - b) > epsilon
}

// Lt returns true if a is strictly less than b. Even if a<b would return true
// they could in fact be equal.
func Lt(a, b float32) bool {
	return a < b && Ne(a, b)
}

// Le returns true if a is less than or equal to b. Even if a<b would return
// true they could in fact be equal.
func Le(a, b float32) bool {
	return a < b || Eq(a, b)
}

// Gt returns true if a is strictly greater than b. Even if a>b would return
// true they could in fact be equal.
func Gt(a, b float32) bool {
	return a > b && Ne(a, b)
}

// Ge returns true if a is greater than or equal to b. Even if a>b would return
// true they could in fact be equal.
func Ge(a, b float32) bool {
	return a > b || Eq(a, b)
}

// Ltz returns true if a is strictly less than b.zero.
func Ltz(a float32) bool {
	return a < 0 && !Z(a)
}

// Lez returns true if a is less than or equal to zero.
func Lez(a float32) bool {
	return a < 0 || Z(a)
}

// Gtz returns true if a is strictly greater than zero.
func Gtz(a float32) bool {
	return a > 0 && !Z(a)
}

// Gez returns true if a is greater than or equal to zero.
func Gez(a float32) bool {
	return a > 0 || Z(a)
}

// refz returns true if a is close to zero. This is the reference implementation
// used in testing and benchmarking.
func refz(a float32) bool {
	if math.Abs(a) <= epsilon*math.Max(1, math.Abs(a)) {
		return true
	}
	return false
}

// Z returns true if a is roughly equal to zero.
func Z(a float32) bool {
	if a > 0 {
		return a <= epsilon
	}
	return -a <= epsilon
}

// Nz returns true if a is not roughly equal to zero.
func Nz(a float32) bool {
	if a > 0 {
		return a >= epsilon
	}
	return -a >= epsilon
}
