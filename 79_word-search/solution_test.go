package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。



 示例 1：


输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
"ABCCED"
输出：true


 示例 2：


输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
"SEE"
输出：true


 示例 3：


输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
"ABCB"
输出：false




 提示：


 m == board.length
 n = board[i].length
 1 <= m, n <= 6
 1 <= word.length <= 15
 board 和 word 仅由大小写英文字母组成




 进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？

 Related Topics 深度优先搜索 数组 字符串 回溯 矩阵 👍 1913 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
var found bool

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			found = false
			dfs(board, i, j, word, 0)
			if found {
				return true
			}
		}
	}
	return false
}

// 从 (i, j) 开始向四周搜索，试图匹配 word[p..]
func dfs(board [][]byte, i, j int, word string, p int) {
	if p == len(word) {
		// 整个 word 已经被匹配完，找到了一个答案
		found = true
		return
	}
	if found {
		// 已经找到了一个答案，不用再搜索了
		return
	}
	m, n := len(board), len(board[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}
	if board[i][j] != word[p] {
		return
	}

	// 已经匹配过的字符，我们给它添一个负号作为标记，避免走回头路
	original := board[i][j]
	board[i][j] = '^' - board[i][j]
	// word[p] 被 board[i][j] 匹配，开始向四周搜索 word[p+1..]
	dfs(board, i+1, j, word, p+1)
	dfs(board, i, j+1, word, p+1)
	dfs(board, i-1, j, word, p+1)
	dfs(board, i, j-1, word, p+1)
	board[i][j] = original
}

//leetcode submit region end(Prohibit modification and deletion)

func TestWordSearch(t *testing.T) {
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
