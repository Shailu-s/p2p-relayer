package big_ext

import (
	"fmt"
	"math/big"
)

// CheckAndParseBigInt parses a string into a big.Int.
func CheckAndParseBigInt(value string) (*big.Int, error) {
	if value == "" {
		return big.NewInt(0), nil
	}
	result, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return nil, fmt.Errorf("invalid int string: %s", value)
	}

	return result, nil
}

// ParseBigInt parses a string into a big.Int. The function panics if the string is invalid.
func ParseBigInt(value string) *big.Int {
	result, err := CheckAndParseBigInt(value)
	if err != nil {
		panic(err)
	}

	return result
}

// CheckAndParseBigFloat parses a string into a big.Float.
func CheckAndParseBigFloat(value string) (*big.Float, error) {
	result, ok := new(big.Float).SetString(value)
	if !ok {
		return nil, fmt.Errorf("invalid float string: %s", value)
	}

	return result, nil
}

// ParseBigFloat parses a string into a big.Float, returning 0 if the string is invalid.
func ParseBigFloat(value string) *big.Float {
	result, err := CheckAndParseBigFloat(value)
	if err != nil {
		panic(err)
	}

	return result
}
