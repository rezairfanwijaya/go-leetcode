package containerwithmostwater

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		Name        string
		Input       []int
		Expectation int
	}{
		{
			Name:        "success",
			Input:       []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			Expectation: 49,
		},
		{
			Name:        "success",
			Input:       []int{1, 1},
			Expectation: 1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := MaxAge(testCase.Input)
			assert.Equal(t, testCase.Expectation, actual)
		})
	}

}
