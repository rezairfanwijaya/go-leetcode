package addindexarray_test

import (
	aia "letcode/add-index-array"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	testCases := []struct {
		Name        string
		NumberA     []int
		NumberB     []int
		Expectation int
		WantError   bool
	}{
		{
			Name:        "success without grather than ten",
			NumberA:     []int{1, 2, 3},
			NumberB:     []int{1, 2, 3},
			Expectation: 642,
			WantError:   false,
		}, {
			Name:        "success with grather than ten",
			NumberA:     []int{1, 8, 3},
			NumberB:     []int{1, 2, 3},
			Expectation: 702,
			WantError:   false,
		}, {
			Name:        "error length not same",
			NumberA:     []int{1, 2, 3, 0},
			NumberB:     []int{1, 8, 3},
			Expectation: 0,
			WantError:   true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.WantError {
				actual := aia.AddTwoNumbers(testCase.NumberA, testCase.NumberB)
				assert.Equal(t, testCase.Expectation, actual)
			} else {
				actual := aia.AddTwoNumbers(testCase.NumberA, testCase.NumberB)
				assert.Equal(t, testCase.Expectation, actual)
			}
		})
	}
}

func TestSplitNumber(t *testing.T) {
	testCase := []struct {
		Name     string
		Number   int
		Expected []int
		Error    bool
	}{
		{
			Name:     "success",
			Number:   10,
			Expected: []int{1, 0},
			Error:    false,
		}, {
			Name:     "error",
			Number:   111,
			Expected: []int{1, 1},
			Error:    true,
		},
	}

	for _, v := range testCase {
		t.Run(v.Name, func(t *testing.T) {
			if v.Error {
				actual := aia.SplitNumber(v.Number)
				assert.Equal(t, v.Expected, actual)
			} else {
				actual := aia.SplitNumber(v.Number)
				assert.Equal(t, v.Expected, actual)
			}
		})
	}
}

func TestGenerateResult(t *testing.T) {
	testCase := struct {
		Name     string
		Slice    []int
		Expected int
	}{
		Name:     "success",
		Slice:    []int{1, 2, 3},
		Expected: 321,
	}

	t.Run(testCase.Name, func(t *testing.T) {
		actual := aia.GenerateResult(testCase.Slice)
		assert.Equal(t, testCase.Expected, actual)
	})
}
