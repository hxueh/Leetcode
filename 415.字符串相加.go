/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
import "strconv"

func reverseString(str string) string {
	strLen := len(str)
	s := make([]byte, strLen)
	for i := strLen - 1; i >= 0; i-- {
		s[strLen-1-i] = str[i]
	}
	return string(s)
}

func addStrings(num1 string, num2 string) string {
	num1 = reverseString(num1)
	num2 = reverseString(num2)
	idx, carry := 0, 0
	res := ""
	for idx < len(num1) || idx < len(num2) || carry > 0 {
		n1, n2 := 0, 0
		if idx < len(num1) {
			n1 = int(num1[idx] - '0')
		}
		if idx < len(num2) {
			n2 = int(num2[idx] - '0')
		}
		curRes := n1 + n2 + carry
		res += strconv.Itoa(curRes % 10)
		carry = curRes / 10
		idx++
	}
	return reverseString(res)
}

// Time complexity O(N)
// Space complexity O(1)

// @lc code=end
