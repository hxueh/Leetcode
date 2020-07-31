/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
func helper(board [][]byte, word string, row, col int) bool {
	if len(word) == 0 {
		return true
	}
	if row < 0 || col < 0 || row >= len(board) || col >= len(board[row]) {
		return false
	}
	if board[row][col] != word[0] {
		return false
	}
	temp := board[row][col]
	board[row][col] = '.'
	word = word[1:]
	if helper(board, word, row-1, col) {
		return true
	}
	if helper(board, word, row+1, col) {
		return true
	}
	if helper(board, word, row, col-1) {
		return true
	}
	if helper(board, word, row, col+1) {
		return true
	}
	board[row][col] = temp
	return false
}

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			if helper(board, word, i, j) {
				return true
			}
		}
	}
	return false
}

// @lc code=end
