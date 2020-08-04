/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	classesIn := make([]int, numCourses)
	classesOut := make(map[int][]int)
	res := 0
	for _, preRequest := range prerequisites {
		wanted, require := preRequest[0], preRequest[1]
		classesIn[wanted]++
		classesOut[require] = append(classesOut[require], wanted)
	}
	queue := make([]int, 0)
	for class, inValue := range classesIn {
		if inValue == 0 {
			queue = append(queue, class)
			res++
		}
	}
	for len(queue) > 0 {
		curLesson := queue[0]
		queue = queue[1:]
		for _, class := range classesOut[curLesson] {
			classesIn[class]--
			if classesIn[class] == 0 {
				queue = append(queue, class)
				res++
			}
		}
	}
	return res == numCourses
}

// @lc code=end
