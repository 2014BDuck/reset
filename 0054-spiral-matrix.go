// Given an m x n matrix, return all elements of the matrix in spiral order.

// Example 1:

// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [1,2,3,6,9,8,7,4,5]
// Example 2:

// Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// Output: [1,2,3,4,8,12,11,10,9,5,6,7]

// Constraints:

// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 10
// -100 <= matrix[i][j] <= 100

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	m := len(matrix[0])

	i, j := 0, 0
	count := 0

	k := 0

	result := make([]int, m*n, m*n)

	for count < m*n {

		result[count] = matrix[i][j]
		matrix[i][j] = 101
		count++
		flag := false

		if count >= m*n {
			break
		}

		for {
			switch k % 4 {
			case 0:
				if j+1 >= m || matrix[i][j+1] == 101 {
					k++
					break
				}
				j++
				flag = true
			case 1:
				if i+1 >= n || matrix[i+1][j] == 101 {
					k++
					break
				}
				i++
				flag = true
			case 2:
				if j-1 < 0 || matrix[i][j-1] == 101 {
					k++
					break
				}
				j--
				flag = true
			case 3:
				if i-1 < 0 || matrix[i-1][j] == 101 {
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
	return result
}