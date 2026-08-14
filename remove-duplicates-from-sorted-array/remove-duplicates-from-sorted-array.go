package removeduplicates

func removeDuplicates(nums []int) int {
	if (len(nums) == 0) {
		return 0
	}

	k := 1
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if (nums[i] != prev) {
			nums[k] = nums[i]
			prev = nums[i]
			k++
		}
	}

	return k
}
