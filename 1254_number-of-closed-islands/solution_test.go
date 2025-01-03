package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
二维矩阵 grid 由 0 （土地）和 1 （水）组成。岛是由最大的4个方向连通的 0 组成的群，封闭岛是一个 完全 由1包围（左、上、右、下）的岛。

 请返回 封闭岛屿 的数目。



 示例 1：




输入：grid = [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,
1],[1,1,1,1,1,1,1,0]]
输出：2
解释：
灰色区域的岛屿是封闭岛屿，因为这座岛屿完全被水域包围（即被 1 区域包围）。

 示例 2：




输入：grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
输出：1


 示例 3：


输入：grid = [[1,1,1,1,1,1,1],
             [1,0,0,0,0,0,1],
             [1,0,1,1,1,0,1],
             [1,0,1,0,1,0,1],
             [1,0,1,1,1,0,1],
             [1,0,0,0,0,0,1],
             [1,1,1,1,1,1,1]]
输出：2




 提示：


 1 <= grid.length, grid[0].length <= 100
 0 <= grid[i][j] <=1


 Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 313 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
// 主函数：计算封闭岛屿的数量
func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for j := 0; j < n; j++ {
		// 把靠上边的岛屿淹掉
		dfs(grid, 0, j)
		// 把靠下边的岛屿淹掉
		dfs(grid, m-1, j)
	}
	for i := 0; i < m; i++ {
		// 把靠左边的岛屿淹掉
		dfs(grid, i, 0)
		// 把靠右边的岛屿淹掉
		dfs(grid, i, n-1)
	}
	// 遍历 grid，剩下的岛屿都是封闭岛屿
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

// 从 (i, j) 开始，将与之相邻的陆地都变成海水
func dfs(grid [][]int, i int, j int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}
	if grid[i][j] == 1 {
		// 已经是海水了
		return
	}
	// 将 (i, j) 变成海水
	grid[i][j] = 1
	// 淹没上下左右的陆地
	dfs(grid, i+1, j)
	dfs(grid, i, j+1)
	dfs(grid, i-1, j)
	dfs(grid, i, j-1)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestNumberOfClosedIslands(t *testing.T) {
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
