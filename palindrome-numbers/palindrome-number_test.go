package palindromenumbers_test

import (
	pn "letcode/palindrome-numbers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPlaindrome(t *testing.T) {
	testCases := []struct {
		Name        string
		Number      int
		Expectation bool
	}{
		{
			Name:        "Test 1",
			Number:      121,
			Expectation: true,
		}, {
			Name:        "Test 2",
			Number:      9090,
			Expectation: false,
		}, {
			Name:        "Test 3",
			Number:      00000,
			Expectation: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := pn.IsPalindrome(testCase.Number)
			assert.Equal(t, testCase.Expectation, actual)
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pn.IsPalindrome(100)
	}
}
