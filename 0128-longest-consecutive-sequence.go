// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

// You must write an algorithm that runs in O(n) time.

// Example 1:

// Input: nums = [100,4,200,1,3,2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
// Example 2:

// Input: nums = [0,3,7,2,5,8,4,6,0,1]
// Output: 9

// Constraints:

// 0 <= nums.length <= 105
// -109 <= nums[i] <= 109

func longestConsecutive(nums []int) int {
	m := map[int]bool{}
	for _, n := range nums {
		m[n] = true
	}

	length := 0
	for k := range m {
		l, r := k-1, k+1
		for {
			if _, ok := m[l]; ok {
				delete(m, l)
				l--

			} else {
				break
			}

		}
		for {
			if _, ok := m[r]; ok {
				delete(m, r)

				r++
			} else {
				break
			}

		}
		length = max(length, r-l-1)
	}
	return length
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}