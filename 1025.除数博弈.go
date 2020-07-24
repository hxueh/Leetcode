/*
 * @lc app=leetcode.cn id=1025 lang=golang
 *
 * [1025] 除数博弈
 */

// @lc code=start
func divisorGame(N int) bool {
    if N == 1 {
        return false
    }
    dp := make([]bool, N + 1)
    dp[0], dp[1], dp[2] = false, false, true
    for i := 3; i < N + 1; i++ {
        for j := 1; j < i; j++ {
            if i % j == 0 && dp[i - j] == false {
                dp[i] = true
                break
            }
        }
    }
    return dp[N]
}

// Time complexity O(N2)
// Space complexity O(N)

// @lc code=end
