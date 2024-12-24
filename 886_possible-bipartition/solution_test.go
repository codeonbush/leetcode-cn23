package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
给定一组 n 人（编号为 1, 2, ..., n）， 我们想把每个人分进任意大小的两组。每个人都可能不喜欢其他人，那么他们不应该属于同一组。

 给定整数 n 和数组 dislikes ，其中 dislikes[i] = [ai, bi] ，表示不允许将编号为 ai 和 bi的人归入同一组。当可以用这种
方法将所有人分进两组时，返回 true；否则返回 false。






 示例 1：


输入：n = 4, dislikes = [[1,2],[1,3],[2,4]]
输出：true
解释：group1 [1,4], group2 [2,3]


 示例 2：


输入：n = 3, dislikes = [[1,2],[1,3],[2,3]]
输出：false


 示例 3：


输入：n = 5, dislikes = [[1,2],[2,3],[3,4],[4,5],[1,5]]
输出：false




 提示：


 1 <= n <= 2000
 0 <= dislikes.length <= 10⁴
 dislikes[i].length == 2
 1 <= dislikes[i][j] <= n
 ai < bi
 dislikes 中每一组都 不同




 Related Topics 深度优先搜索 广度优先搜索 并查集 图 👍 416 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func possibleBipartition(n int, dislikes [][]int) bool {
	// 图节点编号从 1 开始
	color := make([]bool, n+1)
	visited := make([]bool, n+1)
	// 转化成邻接表表示图结构
	graph := buildGraph(n, dislikes)

	// 和之前的 traverse 函数完全相同
	var traverse func(graph [][]int, v int, ok bool) bool
	traverse = func(graph [][]int, v int, ok bool) bool {
		if !ok {
			return ok
		}
		visited[v] = true
		for _, w := range graph[v] {
			if !visited[w] {
				color[w] = !color[v]
				ok = traverse(graph, w, ok)
			} else {
				if color[w] == color[v] {
					return false
				}
			}
		}
		return ok
	}

	ok := true
	for v := 1; v <= n; v++ {
		if !visited[v] {
			ok = traverse(graph, v, ok)
			if !ok {
				return ok
			}
		}
	}

	return ok
}

// 建图函数
func buildGraph(n int, dislikes [][]int) [][]int {
	// 图节点编号为 1...n
	graph := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		graph[i] = []int{}
	}
	for _, edge := range dislikes {
		v := edge[1]
		w := edge[0]
		// 「无向图」相当于「双向图」
		// v -> w
		graph[v] = append(graph[v], w)
		// w -> v
		graph[w] = append(graph[w], v)
	}
	return graph
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPossibleBipartition(t *testing.T) {
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
