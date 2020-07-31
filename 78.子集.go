/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{})
	for _, num := range nums {
		resLen := len(res)
		for i := 0; i < resLen; i++ {
			curRes := make([]int, len(res[i])+1)
			copy(curRes, res[i])
			curRes[len(curRes)-1] = num
			res = append(res, curRes)
		}
		fmt.Println(num, res)
	}
	return res
}

// Time complexity O(N*2N)
// Space complexity O(N*2N)

// @lc code=end
