package removeduplicates

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		wantK    int
		wantNums []int // first wantK elements after the call
	}{
		{
			name:     "example 1",
			nums:     []int{1, 1, 2},
			wantK:    2,
			wantNums: []int{1, 2},
		},
		{
			name:     "example 2",
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			wantK:    5,
			wantNums: []int{0, 1, 2, 3, 4},
		},
		{
			name:     "empty",
			nums:     []int{},
			wantK:    0,
			wantNums: []int{},
		},
		{
			name:     "single element",
			nums:     []int{1},
			wantK:    1,
			wantNums: []int{1},
		},
		{
			name:     "all unique",
			nums:     []int{1, 2, 3},
			wantK:    3,
			wantNums: []int{1, 2, 3},
		},
		{
			name:     "all same",
			nums:     []int{7, 7, 7, 7},
			wantK:    1,
			wantNums: []int{7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int(nil), tt.nums...)

			gotK := removeDuplicates(nums)

			require.Equal(t, tt.wantK, gotK, "k")
			require.GreaterOrEqual(t, gotK, 0)
			require.LessOrEqual(t, gotK, len(nums))
			// append делает non-nil копию: nil и []int{} для задачи эквивалентны
			require.Equal(t, tt.wantNums, append([]int{}, nums[:gotK]...), "nums[:k]")
		})
	}
}
