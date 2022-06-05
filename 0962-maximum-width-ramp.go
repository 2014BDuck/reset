// A ramp in an integer array nums is a pair (i, j) for which i < j and nums[i] <= nums[j]. The width of such a ramp is j - i.

// Given an integer array nums, return the maximum width of a ramp in nums. If there is no ramp in nums, return 0.

// Example 1:

// Input: nums = [6,0,8,2,1,5]
// Output: 4
// Explanation: The maximum width ramp is achieved at (i, j) = (1, 5): nums[1] = 0 and nums[5] = 5.
// Example 2:

// Input: nums = [9,8,1,0,1,9,4,0,4,1]
// Output: 7
// Explanation: The maximum width ramp is achieved at (i, j) = (2, 9): nums[2] = 1 and nums[9] = 1.

// Constraints:

// 2 <= nums.length <= 5 * 104
// 0 <= nums[i] <= 5 * 104

func maxWidthRamp(nums []int) int {
    n := len(nums)
    stack := make([]int, 0, n)
    stack = append(stack, 0)
    res := 0
    for i := 1; i < n; i++ {
        if nums[i] < nums[stack[len(stack)-1]] {
            stack = append(stack, i)
        }
    }
    for j := n - 1; j >= 0; j-- {
        for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[j] {
            res = max(res, j-stack[len(stack)-1])
            stack = stack[:len(stack)-1]
        }
    }
    return res

}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}