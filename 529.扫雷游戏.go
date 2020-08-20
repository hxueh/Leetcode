/*
 * @lc app=leetcode.cn id=529 lang=golang
 *
 * [529] 扫雷游戏
 */

// @lc code=start

func getLandmineNum(board [][]byte, row int, col int, rows int, cols int) int {
	res := 0
	if row > 0 {
		if board[row-1][col] == 'M' {
			res++
		}
		if col > 0 && board[row-1][col-1] == 'M' {
			res++
		}
		if col < cols-1 && board[row-1][col+1] == 'M' {
			res++
		}
	}
	if col > 0 && board[row][col-1] == 'M' {
		res++
	}
	if col < cols-1 && board[row][col+1] == 'M' {
		res++
	}
	if row < rows-1 {
		if board[row+1][col] == 'M' {
			res++
		}
		if col > 0 && board[row+1][col-1] == 'M' {
			res++
		}
		if col < cols-1 && board[row+1][col+1] == 'M' {
			res++
		}
	}
	return res
}

func updateBoard(board [][]byte, click []int) [][]byte {
	if len(board) == 0 || len(board[0]) == 0 {
		return board
	}
	rows, cols := len(board), len(board[0])
	row, col := click[0], click[1]
	if row < 0 || row >= rows || col < 0 || col >= cols {
		return board
	}
	if board[row][col] == 'M' {
		board[row][col] = 'X'
		return board
	} else if board[row][col] == 'E' {
		if landmineNum := getLandmineNum(board, row, col, rows, cols); landmineNum == 0 {
			board[row][col] = 'B'
			updateBoard(board, []int{row + 1, col})
			updateBoard(board, []int{row - 1, col})
			updateBoard(board, []int{row, col + 1})
			updateBoard(board, []int{row, col - 1})
			updateBoard(board, []int{row + 1, col + 1})
			updateBoard(board, []int{row + 1, col - 1})
			updateBoard(board, []int{row - 1, col + 1})
			updateBoard(board, []int{row - 1, col - 1})
		} else {
			board[row][col] = byte(int('0') + landmineNum)
		}
	}
	return board
}

// @lc code=end
