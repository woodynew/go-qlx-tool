package utils

import (
	"math/big"
)

func BigFloatToString(value float64) string {
	oldNum := value
	newNum := big.NewRat(1, 1)
	newNum.SetFloat64(oldNum)
	return newNum.FloatString(0)
}
