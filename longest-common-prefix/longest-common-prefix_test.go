package longestcommonprefix_test

import (
	lcf "letcode/longest-common-prefix"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	testCases := []struct {
		Name        string
		Word        []string
		Expectation string
	}{
		{
			Name:        "Test 1",
			Word:        []string{"flower", "flow", "flight"},
			Expectation: "fl",
		}, {
			Name:        "Test 2",
			Word:        []string{"aku", "kamu", "kita"},
			Expectation: "",
		}, {
			Name:        "Test 3",
			Word:        []string{"aku", "aku", "aku"},
			Expectation: "aku",
		}, {
			Name:        "Test 4",
			Word:        []string{"", "", ""},
			Expectation: "",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := lcf.LongestCommonPrefix(testCase.Word)
			assert.Equal(t, testCase.Expectation, actual)
		})
	}
}
