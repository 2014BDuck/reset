// Given an unsorted integer array nums, return the smallest missing positive integer.

// You must implement an algorithm that runs in O(n) time and uses constant extra space.

// Example 1:

// Input: nums = [1,2,0]
// Output: 3
// Example 2:

// Input: nums = [3,4,-1,1]
// Output: 2
// Example 3:

// Input: nums = [7,8,9,11,12]
// Output: 1

// Constraints:

// 1 <= nums.length <= 5 * 105
// -231 <= nums[i] <= 231 - 1

func firstMissingPositive(nums []int) int {
	nums = append(nums, 0)
	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] >= n || nums[i] < 0 {
			nums[i] = 0
		}
	}
	for i := 0; i < n; i++ {
		nums[nums[i]%n] += n
	}
	for i := 1; i < n; i++ {
		if nums[i]/n == 0 {
			return i
		}
	}
	return n
}