/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
    if m <= 0 {
		copy(nums1, nums2)
        return
	}
	for m > 0 && n > 0 {
        if nums1[m - 1] > nums2[n - 1] {
            nums1[m+n-1] = nums1[m-1]
            m--
        } else {
            nums1[m+n-1] = nums2[n-1]
            n--
        }
    }
    if m == 0 {
        for i := 0; i < n; i++ {
            nums1[i] = nums2[i]
        }
    }
}

// Time complexity O(M+N)
// Space complexity O(1)

// @lc code=end
