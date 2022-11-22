package mergetwosortedlists_test

import (
	mtsl "letcode/merge-two-sorted-lists"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCasesMerge = []struct {
	Name        string
	SliceOne    []int
	SliceTwo    []int
	Expectation []int
}{
	{
		Name:        "Test 1",
		SliceOne:    []int{},
		SliceTwo:    []int{},
		Expectation: []int{},
	}, {
		Name:        "Test 2",
		SliceOne:    []int{},
		SliceTwo:    []int{0},
		Expectation: []int{0},
	}, {
		Name:        "Test 3",
		SliceOne:    []int{1, 2, 3},
		SliceTwo:    []int{1, 2, 3},
		Expectation: []int{1, 1, 2, 2, 3, 3},
	}, {
		Name:        "Test 4",
		SliceOne:    []int{},
		SliceTwo:    []int{1, 2},
		Expectation: []int{1, 2},
	},
}

func TestMergeTwoSortedList(t *testing.T) {
	for _, testCase := range testCasesMerge {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := mtsl.MergeTwoSortedList(testCase.SliceOne, testCase.SliceTwo)
			assert.Equal(t, testCase.Expectation, actual)
		})
	}
}

func BenchmarkMergeTwoSortedList(b *testing.B) {
	for _, testCase := range testCasesMerge {
		for i := 0; i < b.N; i++ {
			mtsl.MergeTwoSortedList(testCase.SliceOne, testCase.SliceTwo)
		}
	}

}

var testCaseSort = []struct {
	Name        string
	Slice       []int
	Expectation []int
}{
	{
		Name:        "test 1",
		Slice:       []int{},
		Expectation: []int{},
	}, {
		Name:        "test 2",
		Slice:       []int{1},
		Expectation: []int{1},
	}, {
		Name:        "test 3",
		Slice:       []int{1, 2, 3, 4},
		Expectation: []int{1, 2, 3, 4},
	}, {
		Name:        "test 4",
		Slice:       []int{5, 4, 2, 3, 1},
		Expectation: []int{1, 2, 3, 4, 5},
	},
}

func TestSorted(t *testing.T) {

	for _, testCase := range testCaseSort {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := mtsl.Sorting(testCase.Slice)
			assert.Equal(t, testCase.Expectation, actual)
		})
	}
}

func BenchmarkTestSorted(b *testing.B) {
	for _, testCase := range testCaseSort {
		for i := 0; i < b.N; i++ {
			mtsl.Sorting(testCase.Slice)
		}
	}

}
