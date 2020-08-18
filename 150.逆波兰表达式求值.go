/*
* @lc app=leetcode.cn id=150 lang=golang
*
* [150] 逆波兰表达式求值
 */

// @lc code=start
import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, str := range tokens {
		switch str {
		case "+":
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "-":
			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "*":
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "/":
			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		default:
			num, _ := strconv.Atoi(str)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

// @lc code=end

