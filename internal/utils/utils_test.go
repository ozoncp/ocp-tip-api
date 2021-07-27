package utils

import (
	"github.com/ozoncp/ocp-tip-api/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitTipsByBatches(t *testing.T) {
	fixtures := []struct {
		sourceSlice    []models.Tip
		batchSize      int
		expectedResult [][]models.Tip
	}{
		{
			sourceSlice: []models.Tip{
				{Id: 1, UserId: 1, ProblemId: 1, Text: "tip 1"},
				{Id: 2, UserId: 2, ProblemId: 2, Text: "tip 2"},
				{Id: 3, UserId: 3, ProblemId: 3, Text: "tip 3"},
				{Id: 4, UserId: 4, ProblemId: 4, Text: "tip 4"},
				{Id: 5, UserId: 5, ProblemId: 5, Text: "tip 5"},
				{Id: 6, UserId: 6, ProblemId: 6, Text: "tip 6"},
				{Id: 7, UserId: 7, ProblemId: 7, Text: "tip 7"},
			},
			batchSize: 3,
			expectedResult: [][]models.Tip{
				{
					{Id: 1, UserId: 1, ProblemId: 1, Text: "tip 1"},
					{Id: 2, UserId: 2, ProblemId: 2, Text: "tip 2"},
					{Id: 3, UserId: 3, ProblemId: 3, Text: "tip 3"},
				},
				{
					{Id: 4, UserId: 4, ProblemId: 4, Text: "tip 4"},
					{Id: 5, UserId: 5, ProblemId: 5, Text: "tip 5"},
					{Id: 6, UserId: 6, ProblemId: 6, Text: "tip 6"},
				},
				{
					{Id: 7, UserId: 7, ProblemId: 7, Text: "tip 7"},
				},
			},
		},
		{
			sourceSlice: []models.Tip{
				{Id: 1, UserId: 1, ProblemId: 1, Text: "tip 1"},
				{Id: 2, UserId: 2, ProblemId: 2, Text: "tip 2"},
				{Id: 3, UserId: 3, ProblemId: 3, Text: "tip 3"},
				{Id: 4, UserId: 4, ProblemId: 4, Text: "tip 4"},
				{Id: 5, UserId: 5, ProblemId: 5, Text: "tip 5"},
				{Id: 6, UserId: 6, ProblemId: 6, Text: "tip 6"},
				{Id: 7, UserId: 7, ProblemId: 7, Text: "tip 7"},
				{Id: 8, UserId: 8, ProblemId: 8, Text: "tip 8"},
			},
			batchSize: 3,
			expectedResult: [][]models.Tip{
				{
					{Id: 1, UserId: 1, ProblemId: 1, Text: "tip 1"},
					{Id: 2, UserId: 2, ProblemId: 2, Text: "tip 2"},
					{Id: 3, UserId: 3, ProblemId: 3, Text: "tip 3"},
				},
				{
					{Id: 4, UserId: 4, ProblemId: 4, Text: "tip 4"},
					{Id: 5, UserId: 5, ProblemId: 5, Text: "tip 5"},
					{Id: 6, UserId: 6, ProblemId: 6, Text: "tip 6"},
				},
				{
					{Id: 7, UserId: 7, ProblemId: 7, Text: "tip 7"},
					{Id: 8, UserId: 8, ProblemId: 8, Text: "tip 8"},
				},
			},
		},
		{
			sourceSlice: []models.Tip{
				{Id: 1, UserId: 1, ProblemId: 1, Text: "tip 1"},
				{Id: 2, UserId: 2, ProblemId: 2, Text: "tip 2"},
				{Id: 3, UserId: 3, ProblemId: 3, Text: "tip 3"},
				{Id: 4, UserId: 4, ProblemId: 4, Text: "tip 4"},
				{Id: 5, UserId: 5, ProblemId: 5, Text: "tip 5"},
				{Id: 6, UserId: 6, ProblemId: 6, Text: "tip 6"},
				{Id: 7, UserId: 7, ProblemId: 7, Text: "tip 7"},
				{Id: 8, UserId: 8, ProblemId: 8, Text: "tip 8"},
				{Id: 9, UserId: 9, ProblemId: 9, Text: "tip 9"},
			},
			batchSize: 3,
			expectedResult: [][]models.Tip{
				{
					{Id: 1, UserId: 1, ProblemId: 1, Text: "tip 1"},
					{Id: 2, UserId: 2, ProblemId: 2, Text: "tip 2"},
					{Id: 3, UserId: 3, ProblemId: 3, Text: "tip 3"},
				},
				{
					{Id: 4, UserId: 4, ProblemId: 4, Text: "tip 4"},
					{Id: 5, UserId: 5, ProblemId: 5, Text: "tip 5"},
					{Id: 6, UserId: 6, ProblemId: 6, Text: "tip 6"},
				},
				{
					{Id: 7, UserId: 7, ProblemId: 7, Text: "tip 7"},
					{Id: 8, UserId: 8, ProblemId: 8, Text: "tip 8"},
					{Id: 9, UserId: 9, ProblemId: 9, Text: "tip 9"},
				},
			},
		},
	}

	for _, fixture := range fixtures {
		actualResult := SplitTipsByBatches(fixture.sourceSlice, fixture.batchSize)
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

func TestConvertTipsToMap(t *testing.T) {
	fixtures := []struct {
		tips           []models.Tip
		expectedResult map[uint64]models.Tip
	}{
		{
			tips: []models.Tip{
				{Id: 1, UserId: 1, ProblemId: 1, Text: "tip 1"},
				{Id: 2, UserId: 2, ProblemId: 2, Text: "tip 2"},
				{Id: 3, UserId: 3, ProblemId: 3, Text: "tip 3"},
			},
			expectedResult: map[uint64]models.Tip{
				1: {Id: 1, UserId: 1, ProblemId: 1, Text: "tip 1"},
				2: {Id: 2, UserId: 2, ProblemId: 2, Text: "tip 2"},
				3: {Id: 3, UserId: 3, ProblemId: 3, Text: "tip 3"},
			},
		},
		{
			tips:           []models.Tip{},
			expectedResult: map[uint64]models.Tip{},
		},
	}
	for _, fixture := range fixtures {
		actualResult := ConvertTipsToMap(fixture.tips)
		assert.Equal(t, fixture.expectedResult, actualResult)
	}
}
