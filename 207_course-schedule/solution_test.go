package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。

 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如
果要学习课程 ai 则 必须 先学习课程 bi 。


 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。


 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。



 示例 1：


输入：numCourses = 2, prerequisites = [[1,0]]
输出：true
解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。

 示例 2：


输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。



 提示：


 1 <= numCourses <= 2000
 0 <= prerequisites.length <= 5000
 prerequisites[i].length == 2
 0 <= ai, bi < numCourses
 prerequisites[i] 中的所有课程对 互不相同


 Related Topics 深度优先搜索 广度优先搜索 图 拓扑排序 👍 2065 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 记录一次递归堆栈中的节点
	onPath := make([]bool, numCourses)
	// 记录节点是否被遍历过
	visited := make([]bool, numCourses)
	// 记录图中是否有环
	hasCycle := false

	graph := buildGraph(numCourses, prerequisites)

	for i := 0; i < numCourses; i++ {
		// 遍历图中的所有节点
		traverse(graph, i, &onPath, &visited, &hasCycle)
	}
	// 只要没有循环依赖可以完成所有课程
	return !hasCycle
}

// 图遍历函数，遍历所有路径
func traverse(graph [][]int, s int, onPath *[]bool, visited *[]bool, hasCycle *bool) {
	if *hasCycle {
		// 如果已经找到了环，也不用再遍历了
		return
	}

	if (*onPath)[s] {
		// s 已经在递归路径上，说明成环了
		*hasCycle = true

		return
	}

	if (*visited)[s] {
		// 不用再重复遍历已遍历过的节点
		return
	}

	if *hasCycle {
		// 如果已经找到了环，也不用再遍历了
		return
	}

	// 前序代码位置
	(*visited)[s] = true
	(*onPath)[s] = true
	for _, t := range graph[s] {
		traverse(graph, t, onPath, visited, hasCycle)
	}
	// 后序代码位置
	(*onPath)[s] = false
}

// buildGraph 构建有向图
func buildGraph(numCourses int, prerequisites [][]int) []([]int) {
	// 图中共有 numCourses 个节点
	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = []int{}
	}
	for _, edge := range prerequisites {
		from, to := edge[1], edge[0]
		// 添加一条从 from 指向 to 的有向边
		// 边的方向是「被依赖」关系，即修完课程 from 才能修课程 to
		graph[from] = append(graph[from], to)
	}
	return graph
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCourseSchedule(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{10, 6, 8, 5, 11, 9},
			expected: []int{3, 1, 1, 1, 0, 0},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := longestWPI(test.input)
			fmt.Println(result)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For heights %v, expected %v but got %v", test.input, test.expected, result)
			}
		})
	}
}
