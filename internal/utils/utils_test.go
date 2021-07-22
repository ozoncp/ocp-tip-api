package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitByBatches(t *testing.T) {
	fixtures := []struct {
		sourceSlice    []int
		batchSize      int
		expectedResult [][]int
	}{
		{
			sourceSlice:    []int{1, 2, 3, 4, 5, 6, 7},
			batchSize:      3,
			expectedResult: [][]int{{1, 2, 3}, {4, 5, 6}, {7}},
		},
		{
			sourceSlice:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			batchSize:      3,
			expectedResult: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8}},
		},
		{
			sourceSlice:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			batchSize:      3,
			expectedResult: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
	}

	for _, fixture := range fixtures {
		actualResult := SplitByBatches(fixture.sourceSlice, fixture.batchSize)
		assert.Equal(t, fixture.expectedResult, actualResult)
	}
}

func TestReverseMap(t *testing.T) {
	fixtures := []struct {
		sourceMap      map[string]int
		expectedResult map[int]string
	}{
		{
			sourceMap:      map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			expectedResult: map[int]string{1: "a", 2: "b", 3: "c", 4: "d"},
		},
	}
	for _, fixture := range fixtures {
		actualResult := ReverseMap(fixture.sourceMap)
		assert.Equal(t, fixture.expectedResult, actualResult)
	}
}

func TestFilterSlice(t *testing.T) {
	fixtures := []struct {
		sourceSlice    []int
		excludedValues []int
		expectedResult []int
	}{
		{
			sourceSlice:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			excludedValues: []int{1, 3, 5},
			expectedResult: []int{2, 4, 6, 7, 8, 9, 10},
		},
		{
			sourceSlice:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			excludedValues: []int{11, 12, 13},
			expectedResult: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, fixture := range fixtures {
		actualResult := FilterSlice(fixture.sourceSlice, fixture.excludedValues)
		assert.Equal(t, fixture.expectedResult, actualResult)
	}
}
