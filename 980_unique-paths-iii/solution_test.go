package leetcode

import(
	"fmt"
	"reflect"
	"testing"
)

/**
在二维网格 grid 上，有 4 种类型的方格： 

 
 1 表示起始方格。且只有一个起始方格。 
 2 表示结束方格，且只有一个结束方格。 
 0 表示我们可以走过的空方格。 
 -1 表示我们无法跨越的障碍。 
 

 返回在四个方向（上、下、左、右）上行走时，从起始方格到结束方格的不同路径的数目。 

 每一个无障碍方格都要通过一次，但是一条路径中不能重复通过同一个方格。 

 

 示例 1： 

 输入：[[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
输出：2
解释：我们有以下两条路径：
1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2) 

 示例 2： 

 输入：[[1,0,0,0],[0,0,0,0],[0,0,0,2]]
输出：4
解释：我们有以下四条路径： 
1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2),(2,3)
2. (0,0),(0,1),(1,1),(1,0),(2,0),(2,1),(2,2),(1,2),(0,2),(0,3),(1,3),(2,3)
3. (0,0),(1,0),(2,0),(2,1),(2,2),(1,2),(1,1),(0,1),(0,2),(0,3),(1,3),(2,3)
4. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2),(2,3) 

 示例 3： 

 输入：[[0,1],[2,0]]
输出：0
解释：
没有一条路能完全穿过每一个空的方格一次。
请注意，起始和结束方格可以位于网格中的任意位置。
 

 

 提示： 

 
 1 <= grid.length * grid[0].length <= 20 
 

 Related Topics 位运算 数组 回溯 矩阵 👍 374 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译。
// 本代码的正确性已通过力扣验证，如有疑问，可以对照 java 代码查看。

func uniquePathsIII(grid [][]int) int {
	var res, visitedCount, totalCount int
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var startI, startJ int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				startI, startJ = i, j
			}
			if grid[i][j] == 1 || grid[i][j] == 0 {
				totalCount++
			}
		}
	}

	dfs(grid, startI, startJ, &res, &visitedCount, totalCount, visited, dirs)
	return res
}

// DFS 算法框架
func dfs(grid [][]int, i, j int, res, visitedCount *int, totalCount int, visited [][]bool, dirs [][]int) {
	m, n := len(grid), len(grid[0])
	// 剪枝，索引越界
	if i < 0 || i >= m || j < 0 || j >= n {
		return
	}
	// 剪枝，跳过起点、障碍物、已访问的格子
	if grid[i][j] == -1 || visited[i][j] {
		return
	}

	// 到达终点
	if grid[i][j] == 2 {
		if *visitedCount == totalCount {
			*res++
		}
		return
	}

	visited[i][j] = true
	*visitedCount++

	for _, dir := range dirs {
		dfs(grid, i+dir[0], j+dir[1], res, visitedCount, totalCount, visited, dirs)
	}

	visited[i][j] = false
	*visitedCount--
}
//leetcode submit region end(Prohibit modification and deletion)


func TestUniquePathsIii(t *testing.T){
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