/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverseString(str string) string {
	res := make([]byte, len(str))
	strLen := len(str)
	for i := strLen - 1; i >= 0; i-- {
		res[strLen-1-i] = str[i]
	}
	return string(res)
}

func addBinary(a string, b string) string {
	a, b = reverseString(a), reverseString(b)
	res := make([]byte, max(len(a), len(b))+1)
	carry := 0
	for i := 0; i < len(res); i++ {
		aNum, bNum := 0, 0
		if i < len(a) {
			aNum = int(a[i] - '0')
		}
		if i < len(b) {
			bNum = int(b[i] - '0')
		}
		curAns := aNum + bNum + carry
		res[i] = byte(curAns%2 + '0')
		carry = curAns / 2
	}
	if res[len(res)-1] == '0' {
		res = res[:len(res)-1]
	}
	if len(res) == 0 {
		return "0"
	}
	return reverseString(string(res))
}

// @lc code=end
