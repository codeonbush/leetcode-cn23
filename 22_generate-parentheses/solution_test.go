package leetcode

import(
	"fmt"
	"reflect"
	"testing"
)

/**
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。 

 

 示例 1： 

 
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
 

 示例 2： 

 
输入：n = 1
输出：["()"]
 

 

 提示： 

 
 1 <= n <= 8 
 

 Related Topics 字符串 动态规划 回溯 👍 3746 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	var track string
	var res []string

	var backtrack func(left, right int)
	backtrack = func(left, right int) {
		// 若左括号剩下的多，说明不合法
		if right < left {
			return
		}
		// 数量小于 0 肯定是不合法的
		if left < 0 || right < 0 {
			return
		}
		// 当所有括号都恰好用完时，得到一个合法的括号组合
		if left == 0 && right == 0 {
			res = append(res, track)
			return
		}

		// 做选择，尝试放一个左括号
		track += "("
		backtrack(left-1, right)
		// 撤销选择
		track = track[:len(track)-1]

		// 做选择，尝试放一个右括号
		track += ")"
		backtrack(left, right-1)
		// 撤销选择
		track = track[:len(track)-1]
	}

	backtrack(n, n)
	return res
}
//leetcode submit region end(Prohibit modification and deletion)


func TestGenerateParentheses(t *testing.T){
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