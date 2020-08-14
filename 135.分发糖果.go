/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sum(list []int) int {
	res := 0
	for _, num := range list {
		res += num
	}
	return res
}

func candy(ratings []int) int {
	ratings = append(ratings, math.MaxInt32)
	resList, stack := make([]int, len(ratings)), make([]int, 1)
	for i := 1; i < len(ratings); i++ {
		curRating, lastRating := ratings[i], ratings[stack[len(stack)-1]]
		if lastRating > curRating {
			stack = append(stack, i)
		} else {
			candyNum := 1
			if stack[0] > 0 && ratings[stack[0]-1] < ratings[stack[0]] {
				candyNum = resList[stack[0]-1] + 1
			}
			candyNum = max(candyNum, len(stack))
			for j := 0; j < len(stack); j++ {
				resList[stack[j]] = candyNum
				candyNum = len(stack) - j - 1
			}
			stack = []int{i}
		}
	}
	return sum(resList)
}

// @lc code=end
