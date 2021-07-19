package utils

import (
	"reflect"
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

	for i, fixture := range fixtures {
		actualResult := SplitByBatches(fixture.sourceSlice, fixture.batchSize)
		if !reflect.DeepEqual(actualResult, fixture.expectedResult) {
			t.Fatalf("Test case %v fail. Got: %v, expected: %v", i+1, actualResult, fixture.expectedResult)
		}
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
	for i, fixture := range fixtures {
		actualResult := ReverseMap(fixture.sourceMap)
		if !reflect.DeepEqual(actualResult, fixture.expectedResult) {
			t.Fatalf("Test case %v fail. Got: %v, expected: %v", i+1, actualResult, fixture.expectedResult)
		}
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
	for i, fixture := range fixtures {
		actualResult := FilterSlice(fixture.sourceSlice, fixture.excludedValues)
		if !reflect.DeepEqual(actualResult, fixture.expectedResult) {
			t.Fatalf("Test case %v fail. Got: %v, expected: %v", i+1, actualResult, fixture.expectedResult)
		}
	}
}
