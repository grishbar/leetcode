/*
Majority Element (LeetCode 169)

https://leetcode.com/problems/majority-element/description

Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times.
You may assume that the majority element always exists in the array.

Example 1:
  Input: nums = [3,2,3]
  Output: 3

Example 2:
  Input: nums = [2,2,1,1,1,2,2]
  Output: 2

Constraints:
  n == nums.length
  1 <= n <= 5 * 10^4
  -10^9 <= nums[i] <= 10^9
  The input is generated such that a majority element will exist in the array.

Follow-up: Could you solve the problem in linear time and in O(1) space?
*/
package majorityelement

func majorityElement(nums []int) int {
  numsCounts := make(map[int]int)

  for i := 0; i < len(nums); i++ {
    numsCounts[nums[i]] = numsCounts[nums[i]] + 1

    if (numsCounts[nums[i]] > len(nums) / 2) {
      return nums[i]
    }
  }

	return nums[0]
}
