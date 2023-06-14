package reverseinteger

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	// is negatif ?
	isNegatif := false
	if x < 0 {
		x = int(math.Abs(float64(x)))
		isNegatif = true
	}

	// to string
	stringX := strconv.Itoa(x)

	// result string
	var stringResult string

	// result number
	var result int

	for i := len(stringX) - 1; i >= 0; i-- {
		stringResult += string(stringX[i])
	}

	result, _ = strconv.Atoi(stringResult)
	if result < math.MinInt32 || result > math.MaxInt32 {
		return 0
	}

	if isNegatif {
		return result * -1
	}

	return result
}
