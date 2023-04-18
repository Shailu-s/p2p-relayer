package big_ext

import "math/big"

func ToFloat64(value *big.Float) float64 {
	f64, _ := value.Float64()
	return f64
}
