package leetcode

import(
	"fmt"
	"reflect"
	"testing"
)

/**
编写一个程序，通过填充空格来解决数独问题。 

 数独的解法需 遵循如下规则： 

 
 数字 1-9 在每一行只能出现一次。 
 数字 1-9 在每一列只能出现一次。 
 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图） 
 

 数独部分空格内已填入了数字，空白格用 '.' 表示。 

 

 
 
 
 示例 1： 
 
 
输入：board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",
".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".",
"3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],
[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".
",".",".",".","8",".",".","7","9"]]
输出：[["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],
["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4
","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6
","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5
","2","8","6","1","7","9"]]
解释：输入的数独如上图所示，唯一有效的解决方案如下所示：
 
 
 
 



 

 提示： 

 
 board.length == 9 
 board[i].length == 9 
 board[i][j] 是一位数字或者 '.' 
 题目数据 保证 输入数独仅有一个解 
 

 Related Topics 数组 哈希表 回溯 矩阵 👍 1890 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func solveSudoku(board [][]byte) {

	// 标记是否已经找到可行解
	found := false


	// 路径：board 中小于 index 的位置所填的数字
	// 选择列表：数字 1~9
	// 结束条件：整个 board 都填满数字
	var backtrack func(index int)
	backtrack = func(index int) {
		if found {
			// 已经找到一个可行解，立即结束
			return
		}

		m, n := 9, 9
		i, j := index/n, index%n
		if index == m*n {
			// 找到一个可行解，触发 base case
			found = true
			return
		}

		if board[i][j] != '.' {
			// 如果有预设数字，不用我们穷举
			backtrack(index + 1)
			return
		}

		for ch := byte('1'); ch <= '9'; ch++ {
			// 剪枝：如果遇到不合法的数字，就跳过
			if !isValid(board, i, j, ch) {
				continue
			}

			// 做选择
			board[i][j] = ch

			backtrack(index + 1)
			if found {
				// 如果找到一个可行解，立即结束
				// 不要撤销选择，否则 board[i][j] 会被重置为 '.'
				return
			}

			// 撤销选择
			board[i][j] = '.'
		}
	}

	backtrack(0)
}

// 判断是否可以在 (r, c) 位置放置数字 num
func isValid(board [][]byte, r, c int, num byte) bool {
	for i := 0; i < 9; i++ {
		// 判断行是否存在重复
		if board[r][i] == num {
			return false
		}
		// 判断列是否存在重复
		if board[i][c] == num {
			return false
		}
		// 判断 3 x 3 方框是否存在重复
		if board[(r/3)*3+i/3][(c/3)*3+i%3] == num {
			return false
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)


func TestSudokuSolver(t *testing.T){
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