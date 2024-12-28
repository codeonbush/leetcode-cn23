package leetcode

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

/**
n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。

 给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。





 示例 1：


输入：n = 4
输出：2
解释：如上图所示，4 皇后问题存在两个不同的解法。




 示例 2：


输入：n = 1
输出：1




 提示：


 1 <= n <= 9


 Related Topics 回溯 👍 557 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func totalNQueens(n int) int {
	// 仅仅记录合法结果的数量
	res := 0

	// 每个字符串代表一行，字符串列表代表一个棋盘
	// '.' 表示空，'Q' 表示皇后，初始化空棋盘
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}

	// 路径：board 中小于 row 的那些行都已经成功放置了皇后
	// 选择列表：第 row 行的所有列都是放置皇后的选择
	// 结束条件：row 超过 board 的最后一行
	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			// 找到一个合法结果
			res++
			return
		}

		for col := 0; col < n; col++ {
			// 剪枝，排除不合法选择
			if !isValid(board, row, col) {
				continue
			}

			// 做选择
			rowChars := []rune(board[row])
			rowChars[col] = 'Q'
			board[row] = string(rowChars)

			// 进入下一行决策
			backtrack(row + 1)

			// 撤销选择
			rowChars[col] = '.'
			board[row] = string(rowChars)
		}
	}

	backtrack(0)
	return res
}

// 是否可以在 board[row][col] 放置皇后？
func isValid(board []string, row, col int) bool {
	// 因为我们是从上往下放置皇后的，
	// 所以只需要检查上方是否有皇后互相冲突，
	// 不需要检查下方

	n := len(board)

	// 检查列是否有皇后互相冲突
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// 检查右上方是否有皇后互相冲突
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// 检查左上方是否有皇后互相冲突
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestNQueensIi(t *testing.T) {
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
