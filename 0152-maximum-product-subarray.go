// Given an integer array nums, find a contiguous non-empty subarray within the array that has the largest product, and return the product.

// The test cases are generated so that the answer will fit in a 32-bit integer.

// A subarray is a contiguous subsequence of the array.

// Example 1:

// Input: nums = [2,3,-2,4]
// Output: 6
// Explanation: [2,3] has the largest product 6.
// Example 2:

// Input: nums = [-2,0,-1]
// Output: 0
// Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

// Constraints:

// 1 <= nums.length <= 2 * 104
// -10 <= nums[i] <= 10
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

func maxProduct(nums []int) int {
	n := len(nums)

	revNums := make([]int, n, n)
	for i := 0; i < n; i++ {
		j := n - i - 1
		revNums[j] = nums[i]
	}

	for i := 1; i < n; i++ {
		if nums[i] != 0 && nums[i-1] != 0 {
			nums[i] = nums[i] * nums[i-1]
		}

		if revNums[i] != 0 && revNums[i-1] != 0 {
			revNums[i] = revNums[i] * revNums[i-1]
		}
	}

	result := math.MinInt
	for _, v := range nums {
		if v > result {
			result = v
		}
	}

	for _, v := range revNums {
		if v > result {
			result = v
		}
	}
	return result
}
