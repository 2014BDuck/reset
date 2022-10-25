func countBattleships(board [][]byte) int {
    n := len(board)
    m := len(board[0])
    result := 0
    for i := 0; i <n; i++ {
        for j := 0; j < m; j++ {
            if board[i][j] == 'X' {
                dfs(board, i, j)
                result++
            }
        }
    }
    return result
}

func dfs(board [][]byte, i, j int) {
    n := len(board)
    m := len(board[0])
    if i < 0 || i >= n || j <0 || j >= m {
        return
    }
    if board[i][j] == 'X' {
        board[i][j] = '.'
        dfs(board, i+1, j)
        dfs(board, i-1, j)
        dfs(board, i, j+1)
        dfs(board, i, j-1)
    }
    return
}
