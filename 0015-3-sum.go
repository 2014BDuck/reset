// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

// Notice that the solution set must not contain duplicate triplets.

// Example 1:

// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
// Example 2:

// Input: nums = []
// Output: []
// Example 3:

// Input: nums = [0]
// Output: []

// Constraints:

// 0 <= nums.length <= 3000
// -105 <= nums[i] <= 105

func threeSum(nums []int) [][]int {

	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for r > l {
			total := nums[i] + nums[l] + nums[r]
			if total == 0 {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for r > l && nums[l] == nums[l-1] {
					l++
				}
				for r > l && nums[r] == nums[r+1] {
					r--
				}
			} else if total > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return result
}