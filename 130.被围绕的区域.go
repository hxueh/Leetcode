/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */

// @lc code=start
func dfs(board [][]byte, row int, col int) {
	if row < 0 || row > len(board)-1 || col < 0 || col > len(board[row])-1 {
		return
	}
	if board[row][col] == 'O' {
		board[row][col] = '.'
		dfs(board, row+1, col)
		dfs(board, row-1, col)
		dfs(board, row, col+1)
		dfs(board, row, col-1)
	}
}

func solve(board [][]byte) {
	if len(board) < 3 || len(board[0]) < 3 {
		return
	}
	for col := 0; col < len(board[0]); col++ {
		dfs(board, 0, col)
		dfs(board, len(board)-1, col)
	}
	for row := 0; row < len(board); row++ {
		dfs(board, row, 0)
		dfs(board, row, len(board[row])-1)
	}
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if board[row][col] == '.' {
				board[row][col] = 'O'
			} else {
				board[row][col] = 'X'
			}
		}
	}
}

// @lc code=end
