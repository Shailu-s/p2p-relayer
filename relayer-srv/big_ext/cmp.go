package big_ext

import "math/big"

// Equals returns true if `x` and `y` are equal.
func Equals(x, y *big.Int) bool {
	return x.Cmp(y) == 0
}

// GreaterThan returns true if `x` is greater than `y`.
func GreaterThan(x, y *big.Int) bool {
	return x.Cmp(y) > 0
}

// GreaterThanOrEqual returns true if `x` is greater than or equal to `y`.
func GreaterThanOrEqual(x, y *big.Int) bool {
	return x.Cmp(y) >= 0
}

// LessThan returns true if `x` is less than `y`.
func LessThan(x, y *big.Int) bool {
	return x.Cmp(y) < 0
}

// LessThanOrEqual returns true if `x` is less than or equal to `y`.
func LessThanOrEqual(x, y *big.Int) bool {
	return x.Cmp(y) <= 0
}

// FloatEquals returns true if `x` and `y` are equal.
func FloatEquals(x, y *big.Float) bool {
	return x.Cmp(y) == 0
}

// FloatGreaterThan returns true if `x` is greater than `y`.
func FloatGreaterThan(x, y *big.Float) bool {
	return x.Cmp(y) > 0
}

// FloatGreaterThanOrEqual returns true if `x` is greater than or equal to `y`.
func FloatGreaterThanOrEqual(x, y *big.Float) bool {
	return x.Cmp(y) >= 0
}

// FloatLessThan returns true if `x` is less than `y`.
func FloatLessThan(x, y *big.Float) bool {
	return x.Cmp(y) < 0
}

// FloatLessThanOrEqual returns true if `x` is less than or equal to `y`.
func FloatLessThanOrEqual(x, y *big.Float) bool {
	return x.Cmp(y) <= 0
}
