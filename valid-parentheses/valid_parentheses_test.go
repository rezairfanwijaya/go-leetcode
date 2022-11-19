package validparentheses_test

import (
	vp "letcode/valid-parentheses"
	"testing"

	"github.com/stretchr/testify/assert"
)

var TestCases = []struct {
	Name        string
	Parentheses string
	Expectation bool
}{
	{
		Name:        "()",
		Parentheses: "()",
		Expectation: true,
	}, {
		Name:        "(){}[]",
		Parentheses: "(){}[]",
		Expectation: true,
	}, {
		Name:        "({[]})",
		Parentheses: "({[]})",
		Expectation: true,
	}, {
		Name:        "{}[()]",
		Parentheses: "{}[()]",
		Expectation: true,
	}, {
		Name:        "({[]})([]){}",
		Parentheses: "({[]})([]){}",
		Expectation: true,
	}, {
		Name:        "(}){",
		Parentheses: "(}){",
		Expectation: false,
	}, {
		Name:        "(",
		Parentheses: "(",
		Expectation: false,
	}, {
		Name:        "",
		Parentheses: "",
		Expectation: false,
	}, {
		Name:        "[[]",
		Parentheses: "[[]",
		Expectation: false,
	}, {
		Name:        "[[}",
		Parentheses: "[[}",
		Expectation: false,
	}, {
		Name:        "({}){]()",
		Parentheses: "({}){]()",
		Expectation: false,
	}, {
		Name:        "(}",
		Parentheses: "(}",
		Expectation: false,
	},
}

func TestIsValid(t *testing.T) {
	for _, testCase := range TestCases {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := vp.IsValid(testCase.Parentheses)
			assert.Equal(t, testCase.Expectation, actual)
		})
	}
}

func BenchmarkIsValid(b *testing.B) {
	for _, testCase := range TestCases {
		b.Run(testCase.Name, func(b *testing.B) {
			vp.IsValid(testCase.Parentheses)
		})
	}
}
