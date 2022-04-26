// Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.

// Example 1:

// Input: nums = [1,1,1], k = 2
// Output: 2
// Example 2:

// Input: nums = [1,2,3], k = 3
// Output: 2

// Constraints:

// 1 <= nums.length <= 2 * 104
// -1000 <= nums[i] <= 1000
// -107 <= k <= 107

func subarraySum(nums []int, k int) int {
    n := len(nums)
    if n == 1 && nums[0] != k {
        return 0
    }

    preSum := make([]int, n, n)
    preSum[0] = nums[0]

    for i := 1; i < n; i++ {
        preSum[i] = preSum[i-1] + nums[i]
    }

    m := map[int]int{0: 1}
    count := 0
    for _, v := range preSum {
        if t, ok := m[v-k]; ok {
            count += t
        }
        m[v] = m[v] + 1
    }

    return count
}