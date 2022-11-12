package romantointeger_test

import (
	rti "letcode/roman-to-integer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	testCases := []struct {
		Name        string
		Symbol      string
		Expectation int
	}{
		{
			Name:        "Test 1",
			Symbol:      "III",
			Expectation: 3,
		}, {
			Name:        "Test 2",
			Symbol:      "LVIII",
			Expectation: 58,
		}, {
			Name:        "Test 3",
			Symbol:      "MCMXCIV",
			Expectation: 1994,
		}, {
			Name:        "Test 4",
			Symbol:      "mcmxciv",
			Expectation: 1994,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := rti.RomanToInt(testCase.Symbol)
			assert.Equal(t, testCase.Expectation, actual)
		})
	}
}

func BenchmarkRomanToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rti.RomanToInt("III")
	}
}
