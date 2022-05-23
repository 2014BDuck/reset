// Given an m x n binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.

// Example 1:

// Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
// Output: 4
// Example 2:

// Input: matrix = [["0","1"],["1","0"]]
// Output: 1
// Example 3:

// Input: matrix = [["0"]]
// Output: 0

// Constraints:

// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 300
// matrix[i][j] is '0' or '1'.

func maximalSquare(matrix [][]byte) int {
	n := len(matrix)
	m := len(matrix[0])
	dp := make([][]int, n, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m, m)
	}

	result := 0
	for i := 0; i < n; i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			result = 1
		}
	}

	for i := 0; i < m; i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
			result = 1
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == '1' {
				width := min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1])
				dp[i][j] = width + 1
				if dp[i][j] > result {
					result = dp[i][j]
				}
			}
		}
	}
	return result * result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}