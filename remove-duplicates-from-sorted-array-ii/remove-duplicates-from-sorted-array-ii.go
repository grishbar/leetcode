package removeduplicatesii

func removeDuplicates(nums []int) int {
	if (len(nums) == 0) {
		return 0
	}
	prev := nums[0]
	prevAmount := 1
	k := 1

	for i := 1; i < len(nums); i++ {
		if (nums[i] != prev) {
			prev = nums[i]
			k++
			prevAmount = 1
		} else {
			
		}
	}

	return k
}
