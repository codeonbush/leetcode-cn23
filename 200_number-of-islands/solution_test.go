package leetcode

import(
	"fmt"
	"reflect"
	"testing"
)

/**
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。 

 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。 

 此外，你可以假设该网格的四条边均被水包围。 

 

 示例 1： 

 
输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1
 

 示例 2： 

 
输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3
 

 

 提示： 

 
 m == grid.length 
 n == grid[i].length 
 1 <= m, n <= 300 
 grid[i][j] 的值为 '0' 或 '1' 
 

 Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 2664 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) int {
	res := 0
	m, n := len(grid), len(grid[0])
	// 遍历 grid
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				// 每发现一个岛屿，岛屿数量加一
				res++
				// 然后使用 DFS 将岛屿淹了
				dfs(grid, i, j)
			}
		}
	}
	return res
}

// 从 (i, j) 开始，将与之相邻的陆地都变成海水
func dfs(grid [][]byte, i, j int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		// 超出索引边界
		return
	}
	if grid[i][j] == '0' {
		// 已经是海水了
		return
	}
	// 将 (i, j) 变成海水
	grid[i][j] = '0'
	// 淹没上下左右的陆地
	dfs(grid, i + 1, j)
	dfs(grid, i, j + 1)
	dfs(grid, i - 1, j)
	dfs(grid, i, j - 1)
}
//leetcode submit region end(Prohibit modification and deletion)


func TestNumberOfIslands(t *testing.T){
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