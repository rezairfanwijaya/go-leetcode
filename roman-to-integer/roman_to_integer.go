package romantointeger

import (
	"strings"
)

func RomanToInt(s string) int {
	romanInt := 0

	// uppercase input
	stringUpperCases := strings.ToUpper(s)

	// symbol romawi
	symbol := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// loop stringUpperCase and match with symbol
	index := 0
	for index < len(stringUpperCases) {
		if index+1 < len(stringUpperCases) {
			if symbol[stringUpperCases[index]] < symbol[stringUpperCases[index+1]] {
				romanInt += symbol[stringUpperCases[index+1]] - symbol[stringUpperCases[index]]
				index += 2
				continue
			} else {
				romanInt += symbol[stringUpperCases[index]]
			}
		} else {
			romanInt += symbol[stringUpperCases[index]]
		}

		index += 1
	}

	return romanInt
}
