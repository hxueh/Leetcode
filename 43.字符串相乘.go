/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start

func add(num1 string, num2 string) string {
	carry := 0
	idx1, idx2 := len(num1)-1, len(num2)-1
	res := ""
	for idx1 >= 0 || idx2 >= 0 || carry > 0 {
		n1, n2 := 0, 0
		if idx1 >= 0 {
			n1 = int(num1[idx1] - '0')
		}
		if idx2 >= 0 {
			n2 = int(num2[idx2] - '0')
		}
		curRes := n1 + n2 + carry
		carry = curRes / 10
		res = string([]byte{byte(curRes%10 + 48)}) + res
		idx1--
		idx2--
	}
	return res
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res, pos := "0", 0
	for i := len(num1) - 1; i >= 0; i-- {
		num1Digit := int(num1[i] - '0')
		curRes, carry := "", 0
		for j := len(num2) - 1; j >= 0; j-- {
			num2Digit := int(num2[j] - '0')
			mul := num1Digit*num2Digit + carry
			curRes = string([]byte{byte(mul%10 + 48)}) + curRes
			carry = mul / 10
		}
		if carry > 0 {
			curRes = string([]byte{byte(carry + 48)}) + curRes
		}
		zeroNumBytes := make([]byte, pos)
		for k := 0; k < pos; k++ {
			zeroNumBytes[k] = '0'
		}
		res = add(res, curRes+string(zeroNumBytes))
		pos++
	}
	return res
}

// @lc code=end
