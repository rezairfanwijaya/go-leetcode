package addtwonumbers_test

import (
	atn "letcode/add-two-numbers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	testCases := []struct {
		Name        string
		Num         []int
		Target      int
		Expextation []int
	}{
		{
			Name:        "Test 1",
			Num:         []int{2, 7, 11, 15},
			Target:      9,
			Expextation: []int{0, 1},
		}, {
			Name:        "Test 2",
			Num:         []int{3, 2, 4},
			Target:      6,
			Expextation: []int{1, 2},
		}, {
			Name:        "Test 3",
			Num:         []int{3, 3},
			Target:      6,
			Expextation: []int{0, 1},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := atn.AddtwoNumber(testCase.Num, testCase.Target)
			assert.Equal(t, testCase.Expextation, actual)
		})
	}
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	num := []int{3, 3}
	target := 6
	for i := 0; i < b.N; i++ {
		atn.AddtwoNumber(num, target)
	}
}
