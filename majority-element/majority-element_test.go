package majorityelement

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{3, 2, 3},
			want: 3,
		},
		{
			name: "example 2",
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
		{
			name: "single element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "two same",
			nums: []int{1, 1},
			want: 1,
		},
		{
			name: "all same",
			nums: []int{7, 7, 7, 7},
			want: 7,
		},
		{
			name: "majority first",
			nums: []int{2, 2, 1},
			want: 2,
		},
		{
			name: "majority last",
			nums: []int{1, 2, 2},
			want: 2,
		},
		{
			name: "negative majority",
			nums: []int{-1, -1, 2},
			want: -1,
		},
		{
			name: "min value majority",
			nums: []int{-1_000_000_000, 1, -1_000_000_000},
			want: -1_000_000_000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int(nil), tt.nums...)

			got := majorityElement(nums)

			require.Equal(t, tt.want, got)
		})
	}
}
