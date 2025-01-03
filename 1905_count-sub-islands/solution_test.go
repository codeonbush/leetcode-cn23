package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
给你两个 m x n 的二进制矩阵 grid1 和 grid2 ，它们只包含 0 （表示水域）和 1 （表示陆地）。一个 岛屿 是由 四个方向 （水平或者竖直）
上相邻的 1 组成的区域。任何矩阵以外的区域都视为水域。

 如果 grid2 的一个岛屿，被 grid1 的一个岛屿 完全 包含，也就是说 grid2 中该岛屿的每一个格子都被 grid1 中同一个岛屿完全包含，那么我
们称 grid2 中的这个岛屿为 子岛屿 。

 请你返回 grid2 中 子岛屿 的 数目 。



 示例 1：
 输入：grid1 = [[1,1,1,0,0],[0,1,1,1,1],[0,0,0,0,0],[1,0,0,0,0],[1,1,0,1,1]], grid2
 = [[1,1,1,0,0],[0,0,1,1,1],[0,1,0,0,0],[1,0,1,1,0],[0,1,0,1,0]]
输出：3
解释：如上图所示，左边为 grid1 ，右边为 grid2 。
grid2 中标红的 1 区域是子岛屿，总共有 3 个子岛屿。


 示例 2：
 输入：grid1 = [[1,0,1,0,1],[1,1,1,1,1],[0,0,0,0,0],[1,1,1,1,1],[1,0,1,0,1]], grid2
 = [[0,0,0,0,0],[1,1,1,1,1],[0,1,0,1,0],[0,1,0,1,0],[1,0,0,0,1]]
输出：2
解释：如上图所示，左边为 grid1 ，右边为 grid2 。
grid2 中标红的 1 区域是子岛屿，总共有 2 个子岛屿。




 提示：


 m == grid1.length == grid2.length
 n == grid1[i].length == grid2[i].length
 1 <= m, n <= 500
 grid1[i][j] 和 grid2[i][j] 都要么是 0 要么是 1 。


 Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 135 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m := len(grid1)
	n := len(grid1[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid1[i][j] == 0 && grid2[i][j] == 1 {
				// 这个岛屿肯定不是子岛，淹掉
				dfs(grid2, i, j)
			}
		}
	}
	// 现在 grid2 中剩下的岛屿都是子岛，计算岛屿数量
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				res++
				dfs(grid2, i, j)
			}
		}
	}
	return res
}

// 从 (i, j) 开始，将与之相邻的陆地都变成海水
func dfs(grid [][]int, i int, j int) {
	m := len(grid)
	n := len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}
	if grid[i][j] == 0 {
		return
	}

	grid[i][j] = 0
	dfs(grid, i+1, j)
	dfs(grid, i, j+1)
	dfs(grid, i-1, j)
	dfs(grid, i, j-1)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCountSubIslands(t *testing.T) {
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
