/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start

func isValidSudoku(board [][]byte) bool {
	everyRow := make([]map[byte]struct{}, 9)
	everyCol := make([]map[byte]struct{}, 9)
	everyBlock := make([]map[byte]struct{}, 9)
	for i := 0; i < 9; i++ {
		everyRow[i] = make(map[byte]struct{})
		everyCol[i] = make(map[byte]struct{})
		everyBlock[i] = make(map[byte]struct{})
	}
	for row := range board {
		for col := range board[row] {
			curByte := board[row][col]
			if curByte == '.' {
				continue
			}
			block := (row / 3) * 3 + col / 3
			curRow, curCol, curBlock := everyRow[row], everyCol[col], everyBlock[block]
			if _, ok := curRow[curByte]; ok {
				return false
			}
			if _, ok := curCol[curByte]; ok {
				return false
			}
			if _, ok := curBlock[curByte]; ok {
				return false
			}
			everyRow[row][curByte], everyCol[col][curByte], everyBlock[block][curByte] = struct{}{}, struct{}{}, struct{}{}
		}
	}
	return true
}
// @lc code=end
