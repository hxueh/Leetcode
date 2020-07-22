/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start

func max(s []int) (res int) {
	res = s[0]
	for _, num := range s {
		if num > res {
			res = num
		}
	}
	return res
}

func longestPalindrome(s string) string {
	res, resLen := "", 0
	for idx := 0; idx < len(s); idx++ {
		lengthOdd := helper(s, idx, idx)
		lengthEven := helper(s, idx, idx+1)
		curMax := max([]int{resLen, lengthEven, lengthOdd})
		if lengthOdd == curMax {
			start := idx - curMax/2
			res = s[start : start+lengthOdd]
		} else if lengthEven == curMax {
			start := idx - (curMax/2 - 1)
			res = s[start : start+lengthEven]
		}
		resLen = curMax
	}
	return res
}

func helper(s string, leftIdx int, rightIdx int) (length int) {
	for leftIdx >= 0 && rightIdx < len(s) && s[leftIdx] == s[rightIdx] {
		leftIdx--
		rightIdx++
	}
	return rightIdx - leftIdx - 1
}

// Time complexity O(N2)
// Space complexity O(1)

// @lc code=end
