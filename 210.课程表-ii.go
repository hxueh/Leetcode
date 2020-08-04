/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 */

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {
	classesIn := make([]int, numCourses)
	classesOut := make(map[int][]int)
	res := make([]int, 0)
	for _, preRequest := range prerequisites {
		wanted, require := preRequest[0], preRequest[1]
		classesIn[wanted]++
		classesOut[require] = append(classesOut[require], wanted)
	}
	queue := make([]int, 0)
	for class, inValue := range classesIn {
		if inValue == 0 {
			queue = append(queue, class)
			res = append(res, class)
		}
	}
	for len(queue) > 0 {
		curLesson := queue[0]
		queue = queue[1:]
		for _, class := range classesOut[curLesson] {
			classesIn[class]--
			if classesIn[class] == 0 {
				queue = append(queue, class)
				res = append(res, class)
			}
		}
	}
	if len(res) != numCourses {
		return []int{}
	}
	return res
}

// @lc code=end
