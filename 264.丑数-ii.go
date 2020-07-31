/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start
func min(list []int) int {
	res := list[0]
	for i := 1; i < len(list); i++ {
		if res > list[i] {
			res = list[i]
		}
	}
	return res
}

func nthUglyNumber(n int) int {
	uglys := make([]int, n)
	uglys[0] = 1
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < n; i++ {
		curRes := min([]int{uglys[i2] * 2, uglys[i3] * 3, uglys[i5] * 5})
		uglys[i] = curRes
		if curRes == uglys[i2]*2 {
			i2++
		}
		if curRes == uglys[i3]*3 {
			i3++
		}
		if curRes == uglys[i5]*5 {
			i5++
		}
	}
	return uglys[n-1]
}

// @lc code=end
