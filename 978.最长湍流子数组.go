/*
 * @lc app=leetcode.cn id=978 lang=golang
 *
 * [978] 最长湍流子数组
 */

// @lc code=start

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxTurbulenceSize(A []int) int {
	if len(A) <= 1 {
		return len(A)
	}
	res, last := 1, 1
	if A[1] != A[0] {
		last = 2
	}
	for i := 2; i < len(A); i++ {
		if (A[i] > A[i-1] && A[i-1] < A[i-2]) || (A[i] < A[i-1] && A[i-1] > A[i-2]) {
			last++
		} else if A[i] == A[i-1] {
			last = 1
		} else {
			last = 2
		}
		res = max(last, res)
	}
	return res
}

// Time complexity O(N)
// Space complexity O(1)

// @lc code=end
