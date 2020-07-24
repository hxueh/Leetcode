package leetcode

import (
	"testing"
)

func compareTwoIntArray(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestTwoSum(t *testing.T) {
	type twoSumTestCase struct {
		arg1 []int
		arg2 int
		ans  []int
	}
	testCases := []twoSumTestCase{
		twoSumTestCase{arg1: []int{2, 7, 11, 15}, arg2: 9, ans: []int{0, 1}},
	}
	for _, testCase := range testCases {
		if ans := twoSum(testCase.arg1, testCase.arg2); !compareTwoIntArray(ans, testCase.ans) {
			t.Errorf("(%v, %d) expected be %v, but got %v", testCase.arg1, testCase.arg2, testCase.ans, ans)
		}
	}
}
