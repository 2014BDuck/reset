// Given a 2D array of characters grid of size m x n, you need to find if there exists any cycle consisting of the same value in grid.

// A cycle is a path of length 4 or more in the grid that starts and ends at the same cell. From a given cell, you can move to one of the cells adjacent to it - in one of the four directions (up, down, left, or right), if it has the same value of the current cell.

// Also, you cannot move to the cell that you visited in your last move. For example, the cycle (1, 1) -> (1, 2) -> (1, 1) is invalid because from (1, 2) we visited (1, 1) which was the last visited cell.

// Return true if any cycle of the same value exists in grid, otherwise, return false.

// Example 1:

// Input: grid = [["a","a","a","a"],["a","b","b","a"],["a","b","b","a"],["a","a","a","a"]]
// Output: true
// Explanation: There are two valid cycles shown in different colors in the image below:

// Example 2:

// Input: grid = [["c","c","c","a"],["c","d","c","c"],["c","c","e","c"],["f","c","c","c"]]
// Output: true
// Explanation: There is only one valid cycle highlighted in the image below:

// Example 3:

// Input: grid = [["a","b","b"],["b","z","b"],["b","b","a"]]
// Output: false

// Constraints:

// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 500
// grid consists only of lowercase English letters.

const (
    ORIGIN = iota
    LEFT
    RIGHT
    UP
    DOWN
)

var visited map[[2]int]bool

func containsCycle(grid [][]byte) bool {
    n := len(grid)
    m := len(grid[0])

    visited = map[[2]int]bool{}
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if visited[[2]int{i, j}] {
                continue
            }
            if dfs(grid, i, j, 'x', ORIGIN) {
                return true
            }
        }
    }
    return false
}

func dfs(grid [][]byte, x, y int, c byte, from int) bool {
    if from == ORIGIN {
        c = grid[x][y]
    } else {

        if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
            return false
        }
        if grid[x][y] == '*' {
            return true
        }
        if grid[x][y] != c {
            return false
        }
    }
    visited[[2]int{x, y}] = true
    grid[x][y] = '*'
    if from != LEFT {
        if dfs(grid, x, y-1, c, RIGHT) {
            return true
        }
    }
    if from != RIGHT {
        if dfs(grid, x, y+1, c, LEFT) {
            return true
        }
    }
    if from != UP {
        if dfs(grid, x-1, y, c, DOWN) {
            return true
        }
    }
    if from != DOWN {
        if dfs(grid, x+1, y, c, UP) {
            return true
        }
    }
    grid[x][y] = c
    return false
}