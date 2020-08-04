/*
 * @lc app=leetcode.cn id=97 lang=golang
 *
 * [97] 交错字符串
 */

// @lc code=start
type tuple struct {
	s1 string
	s2 string
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	curSet, lastSet := make(map[tuple]struct{}), make(map[tuple]struct{})
	lastSet[tuple{s1: s1, s2: s2}] = struct{}{}
	for i := 0; i < len(s3); i++ {
		curChar := s3[i]
		for collection := range lastSet {
			lastS1, lastS2 := collection.s1, collection.s2
			if len(lastS1) > 0 && lastS1[0] == curChar {
				curSet[tuple{s1: lastS1[1:], s2: lastS2}] = struct{}{}
			}
			if len(lastS2) > 0 && lastS2[0] == curChar {
				curSet[tuple{s1: lastS1, s2: lastS2[1:]}] = struct{}{}
			}
		}
		if len(curSet) == 0 {
			return false
		}
		lastSet, curSet = curSet, make(map[tuple]struct{})
	}
	return true
}

// @lc code=end
