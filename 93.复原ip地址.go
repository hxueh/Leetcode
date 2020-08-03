/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原IP地址
 */

// @lc code=start
import "strconv"

func isValid(str string) bool {
	if len(str) > 1 && str[0] == '0' {
		return false
	}
	if numStr, _ := strconv.Atoi(str); numStr > 255 {
		return false
	}
	return true
}

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	for i := 0; i < 3; i++ {
		if len(s)-i > 10 {
			continue
		}
		if len(s)-i < 4 {
			break
		}
		first := s[:i+1]
		if !isValid(first) {
			continue
		}
		for j := i + 1; j < i+4; j++ {
			if len(s)-j > 7 {
				continue
			}
			if len(s)-j < 3 {
				break
			}
			second := s[i+1 : j+1]
			if !isValid(second) {
				continue
			}
			for k := j + 1; k < j+4; k++ {
				if len(s)-k > 4 {
					continue
				}
				if len(s)-k < 2 {
					break
				}
				third := s[j+1 : k+1]
				if !isValid(third) {
					continue
				}
				forth := s[k+1:]
				if !isValid(forth) {
					continue
				}
				res = append(res, first+"."+second+"."+third+"."+forth)
			}
		}
	}
	return res
}

// @lc code=end
