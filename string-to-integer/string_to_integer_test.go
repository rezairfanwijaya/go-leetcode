package stringtointeger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToInteger(t *testing.T) {
	testCases := []struct {
		Name     string
		Input    string
		Expected int
	}{
		{
			Name:     "Success",
			Input:    "123",
			Expected: 123,
		},
		{
			Name:     "Success",
			Input:    "hallo 78567",
			Expected: 0,
		},
		{
			Name:     "Success",
			Input:    "00567 hallo ",
			Expected: 567,
		},
		{
			Name:     "Success",
			Input:    "         -67",
			Expected: -67,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := StringToInteger(testCase.Input)
			assert.Equal(t, testCase.Expected, actual)
		})
	}
}
