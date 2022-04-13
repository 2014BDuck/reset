// Given an m x n grid of characters board and a string word, return true if word exists in the grid.

// The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.

// Example 1:

// Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
// Output: true
// Example 2:

// Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
// Output: true
// Example 3:

// Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
// Output: false

// Constraints:

// m == board.length
// n = board[i].length
// 1 <= m, n <= 6
// 1 <= word.length <= 15
// board and word consists of only lowercase and uppercase English letters.

// Follow up: Could you use search pruning to make your solution faster with a larger board?

func exist(board [][]byte, word string) bool {
    n := len(board)
    m := len(board[0])
    for x := 0; x < n; x++ {
        for y := 0; y < m; y++ {
            if dfs(board, x, y, word, 0) {
                return true
            }
        }
    }
    return false
}

func dfs(board [][]byte, x, y int, word string, pos int) bool {
    n := len(board)
    m := len(board[0])
    if pos == len(word) {
        return true
    }

    if x >= 0 && x < n && y >= 0 && y < m && board[x][y] == word[pos] {
        tmp := board[x][y]
        board[x][y] = byte(0)
        r := dfs(board, x-1, y, word, pos+1) || dfs(board, x+1, y, word, pos+1) || dfs(board, x, y-1, word, pos+1) || dfs(board, x, y+1, word, pos+1)
        board[x][y] = tmp
        return r
    }
    return false
}