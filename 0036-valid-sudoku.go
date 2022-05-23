// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
// Note:

// A Sudoku board (partially filled) could be valid but is not necessarily solvable.
// Only the filled cells need to be validated according to the mentioned rules.

// Example 1:

// Input: board =
// [["5","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
// Output: true
// Example 2:

// Input: board =
// [["8","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
// Output: false
// Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.

// Constraints:

// board.length == 9
// board[i].length == 9
// board[i][j] is a digit 1-9 or '.'.

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if !validRow(board, i) || !validCol(board, i) {
			return false
		}
	}

	for i := 0; i < 9; i = i + 3 {
		for j := 0; j < 9; j = j + 3 {
			if !validGrid(board, i, j) {
				return false
			}
		}
	}
	return true
}

func validRow(board [][]byte, i int) bool {
	m := map[byte]bool{}

	for j := 0; j < 9; j++ {
		if board[i][j] == '.' {
			continue
		}
		_, ok := m[board[i][j]]
		if ok {
			return false
		}
		m[board[i][j]] = true
	}
	return true
}

func validCol(board [][]byte, i int) bool {
	m := map[byte]bool{}

	for j := 0; j < 9; j++ {
		if board[j][i] == '.' {
			continue
		}
		_, ok := m[board[j][i]]
		if ok {
			return false
		}
		m[board[j][i]] = true
	}
	return true
}

func validGrid(board [][]byte, x, y int) bool {
	m := map[byte]bool{}

	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			if board[i][j] == '.' {
				continue
			}
			_, ok := m[board[i][j]]
			if ok {
				return false
			}
			m[board[i][j]] = true
		}
	}
	return true
}