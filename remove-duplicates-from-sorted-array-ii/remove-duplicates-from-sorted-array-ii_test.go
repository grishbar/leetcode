package removeduplicatesii

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Mirrors the LeetCode judge:
//
//	k := removeDuplicates(nums)
//	assert k == len(expectedNums)
//	for i := 0; i < k; i++ { assert nums[i] == expectedNums[i] }
func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name         string
		nums         []int
		expectedNums []int
	}{
		{
			name:         "example 1",
			nums:         []int{1, 1, 1, 2, 2, 3},
			expectedNums: []int{1, 1, 2, 2, 3},
		},
		{
			name:         "example 2",
			nums:         []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			expectedNums: []int{0, 0, 1, 1, 2, 3, 3},
		},
		{
			name:         "single element",
			nums:         []int{1},
			expectedNums: []int{1},
		},
		{
			name:         "two same",
			nums:         []int{1, 1},
			expectedNums: []int{1, 1},
		},
		{
			name:         "two different",
			nums:         []int{1, 2},
			expectedNums: []int{1, 2},
		},
		{
			name:         "all unique",
			nums:         []int{1, 2, 3},
			expectedNums: []int{1, 2, 3},
		},
		{
			name:         "all same more than twice",
			nums:         []int{7, 7, 7, 7},
			expectedNums: []int{7, 7},
		},
		{
			name:         "exactly two of each",
			nums:         []int{1, 1, 2, 2, 3, 3},
			expectedNums: []int{1, 1, 2, 2, 3, 3},
		},
		{
			name:         "triples only",
			nums:         []int{1, 1, 1, 2, 2, 2},
			expectedNums: []int{1, 1, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int(nil), tt.nums...)

			k := removeDuplicates(nums)

			require.Equal(t, len(tt.expectedNums), k, "k == len(expectedNums)")
			for i := 0; i < k; i++ {
				require.Equal(t, tt.expectedNums[i], nums[i], "nums[%d]", i)
			}
		})
	}
}
