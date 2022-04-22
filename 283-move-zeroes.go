// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Note that you must do this in-place without making a copy of the array.

// Example 1:

// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
// Example 2:

// Input: nums = [0]
// Output: [0]

// Constraints:

// 1 <= nums.length <= 104
// -231 <= nums[i] <= 231 - 1

// Follow up: Could you minimize the total number of operations done?

func moveZeroes(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	count := 0
	for i := 0; i < n; i++ {
		nums[i-count] = nums[i]

		if nums[i] == 0 {
			count++
		}
	}

	for i := n - 1; i >= n-count; i-- {
		nums[i] = 0
	}
	return
}