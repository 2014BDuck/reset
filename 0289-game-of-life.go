// According to Wikipedia's article: "The Game of Life, also known simply as Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970."

// The board is made up of an m x n grid of cells, where each cell has an initial state: live (represented by a 1) or dead (represented by a 0). Each cell interacts with its eight neighbors (horizontal, vertical, diagonal) using the following four rules (taken from the above Wikipedia article):

// Any live cell with fewer than two live neighbors dies as if caused by under-population.
// Any live cell with two or three live neighbors lives on to the next generation.
// Any live cell with more than three live neighbors dies, as if by over-population.
// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
// The next state is created by applying the above rules simultaneously to every cell in the current state, where births and deaths occur simultaneously. Given the current state of the m x n grid board, return the next state.

// Example 1:

// Input: board = [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]
// Output: [[0,0,0],[1,0,1],[0,1,1],[0,1,0]]
// Example 2:

// Input: board = [[1,1],[1,0]]
// Output: [[1,1],[1,1]]

// Constraints:

// m == board.length
// n == board[i].length
// 1 <= m, n <= 25
// board[i][j] is 0 or 1.

// Follow up:

// Could you solve it in-place? Remember that the board needs to be updated simultaneously: You cannot update some cells first and then use their updated values to update other cells.
// In this question, we represent the board using a 2D array. In principle, the board is infinite, which would cause problems when the active area encroaches upon the border of the array (i.e., live cells reach the border). How would you address these problems?

func gameOfLife(board [][]int) {
    n := len(board)
    m := len(board[0])

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            c := count(board, i, j)
            if board[i][j] == 0 {
                board[i][j] = 10 + c
            } else {
                board[i][j] = 20 + c
            }
        }
    }

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            v, c := getValAndCount(board[i][j])
            if c < 2 {
                board[i][j] = 0
            } else if v == 1 && (c == 2 || c == 3) {
                board[i][j] = v
            } else if c > 3 {
                board[i][j] = 0
            } else if v == 0 && c == 3 {
                board[i][j] = 1
            } else {
                board[i][j] = v
            }
        }
    }
}

func getValAndCount(a int) (value, count int) {
    if a < 10 {
        return a, 0
    } else if a >= 10 && a < 20 {
        return 0, a - 10
    } else {
        return 1, a - 20
    }
}

func getVal(a int) (value int) {
    if a < 10 {
        return a
    } else if a >= 10 && a < 20 {
        return 0
    } else {
        return 1
    }
}

func count(board [][]int, i, j int) int {
    n := len(board)
    m := len(board[0])

    c := 0
    if i-1 >= 0 {
        c += getVal(board[i-1][j])
        if j-1 >= 0 {
            c += getVal(board[i-1][j-1])
        }
        if j+1 < m {
            c += getVal(board[i-1][j+1])
        }
    }

    if i+1 < n {
        c += getVal(board[i+1][j])
        if j-1 >= 0 {
            c += getVal(board[i+1][j-1])
        }
        if j+1 < m {
            c += getVal(board[i+1][j+1])
        }
    }
    if j-1 >= 0 {
        c += getVal(board[i][j-1])
    }
    if j+1 < m {
        c += getVal(board[i][j+1])
    }
    return c
}