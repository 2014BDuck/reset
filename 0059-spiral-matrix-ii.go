// Given a positive integer n, generate an n x n matrix filled with elements from 1 to n2 in spiral order.

// Example 1:

// Input: n = 3
// Output: [[1,2,3],[8,9,4],[7,6,5]]
// Example 2:

// Input: n = 1
// Output: [[1]]

// Constraints:

// 1 <= n <= 20

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n, n)
	}
	k := 0

	i, j := 0, 0
	c := 1
	for matrix[i][j] == 0 {
		matrix[i][j] = c
		c++

		flag := false
		for {
			if c > n*n {
				break
			}

			switch k % 4 {
			case 0:
				if j+1 >= n || matrix[i][j+1] != 0 {
					k++
					break
				}
				j++
				flag = true
			case 1:
				if i+1 >= n || matrix[i+1][j] != 0 {
					k++
					break
				}
				i++
				flag = true
			case 2:
				if j-1 < 0 || matrix[i][j-1] != 0 {
					k++
					break
				}
				j--
				flag = true
			case 3:
				if i-1 < 0 || matrix[i-1][j] != 0 {
					k++
					break
				}
				i--
				flag = true
			}
			if flag {
				break
			}
		}
	}

	return matrix
}