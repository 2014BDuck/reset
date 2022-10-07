// You are given two integers m and n, which represent the dimensions of a matrix.

// You are also given the head of a linked list of integers.

// Generate an m x n matrix that contains the integers in the linked list presented in spiral order (clockwise), starting from the top-left of the matrix. If there are remaining empty spaces, fill them with -1.

// Return the generated matrix.

// Example 1:

// Input: m = 3, n = 5, head = [3,0,2,6,8,1,7,9,4,2,5,5,0]
// Output: [[3,0,2,6,8],[5,0,-1,-1,1],[5,2,4,9,7]]
// Explanation: The diagram above shows how the values are printed in the matrix.
// Note that the remaining spaces in the matrix are filled with -1.
// Example 2:

// Input: m = 1, n = 4, head = [0,1,2]
// Output: [[0,1,2,-1]]
// Explanation: The diagram above shows how the values are printed from left to right in the matrix.
// The last space in the matrix is set to -1.

// Constraints:

// 1 <= m, n <= 105
// 1 <= m * n <= 105
// The number of nodes in the list is in the range [1, m * n].
// 0 <= Node.val <= 1000

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func spiralMatrix(m int, n int, head *ListNode) [][]int {
    //prepare
    matrix := [][]int{}
    for i := 0; i < m; i++ {
        row := make([]int, n, n)
        for j := 0; j < n; j++ {
            row[j] = -1
        }
        matrix = append(matrix, row)
    }

    // filling
    row, col := 0, 1
    x, y := 0, 0
    for head != nil {
        matrix[x][y] = head.Val
        if x+row < 0 || x+row >= m || y+col < 0 || y+col >= n || matrix[x+row][y+col] != -1 {
            // change direction
            col, row = -row, col
        }
        x, y = x+row, y+col
        head = head.Next
    }

    return matrix
}
