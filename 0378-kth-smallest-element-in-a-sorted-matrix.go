// Given an n x n matrix where each of the rows and columns is sorted in ascending order, return the kth smallest element in the matrix.

// Note that it is the kth smallest element in the sorted order, not the kth distinct element.

// You must find a solution with a memory complexity better than O(n2).

// Example 1:

// Input: matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
// Output: 13
// Explanation: The elements in the matrix are [1,5,9,10,11,12,13,13,15], and the 8th smallest number is 13
// Example 2:

// Input: matrix = [[-5]], k = 1
// Output: -5

// Constraints:

// n == matrix.length == matrix[i].length
// 1 <= n <= 300
// -109 <= matrix[i][j] <= 109
// All the rows and columns of matrix are guaranteed to be sorted in non-decreasing order.
// 1 <= k <= n2

// Follow up:

// Could you solve the problem with a constant memory (i.e., O(1) memory complexity)?
// Could you solve the problem in O(n) time complexity? The solution may be too advanced for an interview but you may find reading this paper fun.

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	m := len(matrix[0])

	countLessOrEqual := func(x int) int {
		cnt := 0
		c := n - 1
		for r := 0; r < m; r++ {
			for c >= 0 && matrix[r][c] > x {
				c -= 1
			}

			cnt += c + 1
		}
		return cnt
	}

	minItem := matrix[0][0]
	maxItem := matrix[n-1][m-1]
	result := -1

	for maxItem >= minItem {
		mid := (maxItem + minItem) / 2
		if countLessOrEqual(mid) >= k {
			result = mid
			maxItem = mid - 1
		} else {
			minItem = mid + 1
		}
	}
	return result
}

