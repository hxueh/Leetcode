package leetcode

import "testing"

func TestTrap(t *testing.T) {
	type trapTestCase struct {
		arg1 []int
		ans  int
	}
	testCases := []trapTestCase{
		trapTestCase{arg1: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, ans: 6},
		trapTestCase{arg1: []int{6, 7, 9, 2, 7, 2, 4, 5, 8, 0}, ans: 20},
		trapTestCase{arg1: []int{3, 3, 4, 8, 4, 3, 4, 4, 7, 10}, ans: 18},
		trapTestCase{arg1: []int{1}, ans: 0},
		trapTestCase{arg1: []int{}, ans: 0},
		trapTestCase{arg1: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, ans: 0},
		trapTestCase{arg1: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, ans: 0},
	}
	for _, testCase := range testCases {
		if ans := trap(testCase.arg1); ans != testCase.ans {
			t.Errorf("%v expected be %d, but got %d", testCase.arg1, testCase.ans, ans)
		}
	}
}
