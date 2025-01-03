package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
给你一个大小为 m x n 的二进制矩阵 grid 。

 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设 grid 的四个边缘都被
0（代表水）包围着。

 岛屿的面积是岛上值为 1 的单元格的数目。

 计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。



 示例 1：


输入：grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,
0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,
0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
输出：6
解释：答案不应该是 11 ，因为岛屿只能包含水平或垂直这四个方向上的 1 。


 示例 2：


输入：grid = [[0,0,0,0,0,0,0,0]]
输出：0




 提示：


 m == grid.length
 n == grid[i].length
 1 <= m, n <= 50
 grid[i][j] 为 0 或 1


 Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 1112 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func maxAreaOfIsland(grid [][]int) int {
	// 记录岛屿的最大面积
	var res int
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				// 淹没岛屿，并更新最大岛屿面积
				res = max(res, dfs(grid, i, j))
			}
		}
	}
	return res
}

// 淹没与 (i, j) 相邻的陆地，并返回淹没的陆地面积
func dfs(grid [][]int, i, j int) int {
	m := len(grid)
	n := len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		// 超出索引边界
		return 0
	}
	if grid[i][j] == 0 {
		// 已经是海水了
		return 0
	}
	// 将 (i, j) 变成海水
	grid[i][j] = 0

	return dfs(grid, i+1, j) +
		dfs(grid, i, j+1) +
		dfs(grid, i-1, j) +
		dfs(grid, i, j-1) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaxAreaOfIsland(t *testing.T) {
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
