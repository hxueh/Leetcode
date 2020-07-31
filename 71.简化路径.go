/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */

// @lc code=start
import (
	"strings"
)

func simplifyPath(path string) string {
	path = path + "/"
	paths := make([]string, 0)
	start := 0
	for end := 0; end < len(path); end++ {
		if path[end] == '/' {
			if start+1 < end {
				paths = append(paths, string(path[start+1:end]))
			}
			start = end
		}
	}
	idx := 0
	for idx < len(paths) {
		switch paths[idx] {
		case ".":
			paths = append(paths[:idx], paths[idx+1:]...)
		case "..":
			if idx < 1 {
				paths = paths[idx+1:]
			} else {
				paths = append(paths[:idx-1], paths[idx+1:]...)
				idx--
			}
		default:
			idx++
		}
	}
	return "/" + strings.Join(paths, "/")
}

// Time complexity O(N)
// Space complexity O(N)

// @lc code=end
