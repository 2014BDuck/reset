// Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

// If target is not found in the array, return [-1, -1].

// You must write an algorithm with O(log n) runtime complexity.

// Example 1:

// Input: nums = [5,7,7,8,8,10], target = 8
// Output: [3,4]
// Example 2:

// Input: nums = [5,7,7,8,8,10], target = 6
// Output: [-1,-1]
// Example 3:

// Input: nums = [], target = 0
// Output: [-1,-1]

// Constraints:

// 0 <= nums.length <= 105
// -109 <= nums[i] <= 109
// nums is a non-decreasing array.
// -109 <= target <= 109

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}

	// edge case
	if len(nums) == 0 {
		return result
	}

	// find left
	l, r := 0, len(nums)-1
	for r > l {
		m := (r + l) / 2
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}

	if nums[l] == target {
		result[0] = l
	} else {
		return result
	}

	// find right
	l, r = 0, len(nums)-1
	for r > l {
		m := (r+l)/2 + 1
		if nums[m] > target {
			r = m - 1
		} else {
			l = m
		}
	}
	result[1] = l
	return result
}