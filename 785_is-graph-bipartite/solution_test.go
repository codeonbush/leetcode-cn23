package leetcode

import(
	"fmt"
	"reflect"
	"testing"
)

/**
存在一个 无向图 ，图中有 n 个节点。其中每个节点都有一个介于 0 到 n - 1 之间的唯一编号。给你一个二维数组 graph ，其中 graph[u] 是
一个节点数组，由节点 u 的邻接节点组成。形式上，对于 graph[u] 中的每个 v ，都存在一条位于节点 u 和节点 v 之间的无向边。该无向图同时具有以下
属性：

 
 不存在自环（graph[u] 不包含 u）。 
 不存在平行边（graph[u] 不包含重复值）。 
 如果 v 在 graph[u] 内，那么 u 也应该在 graph[v] 内（该图是无向图） 
 这个图可能不是连通图，也就是说两个节点 u 和 v 之间可能不存在一条连通彼此的路径。 
 

 二分图 定义：如果能将一个图的节点集合分割成两个独立的子集 A 和 B ，并使图中的每一条边的两个节点一个来自 A 集合，一个来自 B 集合，就将这个图称为 
二分图 。 

 如果图是二分图，返回 true ；否则，返回 false 。 

 

 示例 1： 
 
 
输入：graph = [[1,2,3],[0,2],[0,1,3],[0,2]]
输出：false
解释：不能将节点分割成两个独立的子集，以使每条边都连通一个子集中的一个节点与另一个子集中的一个节点。 

 示例 2： 
 
 
输入：graph = [[1,3],[0,2],[1,3],[0,2]]
输出：true
解释：可以将节点分成两组: {0, 2} 和 {1, 3} 。 

 

 提示： 

 
 graph.length == n 
 1 <= n <= 100 
 0 <= graph[u].length < n 
 0 <= graph[u][i] <= n - 1 
 graph[u] 不会包含 u 
 graph[u] 的所有值 互不相同 
 如果 graph[u] 包含 v，那么 graph[v] 也会包含 u 
 

 Related Topics 深度优先搜索 广度优先搜索 并查集 图 👍 549 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func isBipartite(graph [][]int) bool {
	n := len(graph)
	color := make([]bool, n)
	visited := make([]bool, n)
	ok := true

	var traverse func(int)
	traverse = func(v int) {
		// 如果已经确定不是二分图了，就不用浪费时间再递归遍历了
		if !ok {
			return
		}

		visited[v] = true
		for _, w := range graph[v] {
			if !visited[w] {
				// 相邻节点 w 没有被访问过
				// 那么应该给节点 w 涂上和节点 v 不同的颜色
				color[w] = !color[v]
				// 继续遍历 w
				traverse(w)
			} else {
				// 相邻节点 w 已经被访问过
				// 根据 v 和 w 的颜色判断是否是二分图
				if color[w] == color[v] {
					// 若相同，则此图不是二分图
					ok = false
					return
				}
			}
		}
	}

	// 因为图不一定是联通的，可能存在多个子图
	// 所以要把每个节点都作为起点进行一次遍历
	// 如果发现任何一个子图不是二分图，整幅图都不算二分图
	for v := 0; v < n; v++ {
		if !visited[v] {
			traverse(v)
		}
	}

	return ok
}
//leetcode submit region end(Prohibit modification and deletion)


func TestIsGraphBipartite(t *testing.T){
    tests := []struct {
		input  []int
		expected []int
	}{
		{
			input:  []int{10, 6, 8, 5, 11, 9},
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