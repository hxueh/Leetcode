/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	symbolMap := map[rune]rune{
		')': '(', '}': '{', ']': '[',
	}
	stack := make([]rune, 0)
	for _, char := range s {
		if wanted, ok := symbolMap[char]; !ok {
			stack = append(stack, char)
		} else if len(stack) == 0 {
            return false
        } else {
			if stack[len(stack) - 1] != wanted {
				return false
			} else {
				stack = stack[:len(stack) - 1]
			}
		}
	}
	return len(stack) == 0
}

// Time complexity O(N)
// Space complexity O(N)

// @lc code=end
