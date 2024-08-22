package piscine

import (
	"math/big"
	"strings"
)

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Convert nbr to a decimal number
	decimalValue := big.NewInt(0)
	baseFromInt := big.NewInt(int64(len(baseFrom)))
	for _, digit := range nbr {
		currentDigit := strings.Index(baseFrom, string(digit))
		decimalValue.Mul(decimalValue, baseFromInt)
		decimalValue.Add(decimalValue, big.NewInt(int64(currentDigit)))
	}

	// Convert the decimal number to the target base
	baseToLength := len(baseTo)
	result := ""
	baseToBigInt := big.NewInt(int64(baseToLength))
	zero := big.NewInt(0)

	for decimalValue.Cmp(zero) > 0 {
		remainder := new(big.Int)
		decimalValue.DivMod(decimalValue, baseToBigInt, remainder)
		result = string(baseTo[remainder.Int64()]) + result
	}

	return result
}
