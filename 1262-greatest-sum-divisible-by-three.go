// Given an array nums of integers, we need to find the maximum possible sum of elements of the array such that it is divisible by three.

// Example 1:

// Input: nums = [3,6,5,1,8]
// Output: 18
// Explanation: Pick numbers 3, 6, 1 and 8 their sum is 18 (maximum sum divisible by 3).
// Example 2:

// Input: nums = [4]
// Output: 0
// Explanation: Since 4 is not divisible by 3, do not pick any number.
// Example 3:

// Input: nums = [1,2,3,4,4]
// Output: 12
// Explanation: Pick numbers 1, 3, 4 and 4 their sum is 12 (maximum sum divisible by 3).

// Constraints:

// 1 <= nums.length <= 4 * 10^4
// 1 <= nums[i] <= 10^4

func maxSumDivThree(nums []int) int {
    n := len(nums)
    dp := [3]int{}

    for i := 0; i < n; i++ {
        tmp := dp
        for _, v := range tmp {
            dp[(v+nums[i])%3] = max(dp[(v+nums[i])%3], v+nums[i])
        }
    }
    return dp[0]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}