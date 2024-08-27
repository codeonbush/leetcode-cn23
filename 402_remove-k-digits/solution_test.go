package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
给你一个以字符串表示的非负整数 num 和一个整数 k ，移除这个数中的 k 位数字，使得剩下的数字最小。请你以字符串形式返回这个最小的数字。

 示例 1 ：


输入：num = "1432219", k = 3
输出："1219"
解释：移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219 。


 示例 2 ：


输入：num = "10200", k = 1
输出："200"
解释：移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。


 示例 3 ：


输入：num = "10", k = 2
输出："0"
解释：从原数字移除所有的数字，剩余为空就是 0 。




 提示：


 1 <= k <= num.length <= 10⁵
 num 仅由若干位数字（0 - 9）组成
 除了 0 本身之外，num 不含任何前导零


 Related Topics 栈 贪心 字符串 单调栈 👍 1068 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func removeKdigits(num string, k int) string {
	stk := []rune{}
	for _, c := range num {
		// 单调栈代码模板
		for len(stk) > 0 && c < stk[len(stk)-1] && k > 0 {
			stk = stk[:len(stk)-1]
			k--
		}
		// 防止 0 作为数字的开头
		if len(stk) == 0 && c == '0' {
			continue
		}
		stk = append(stk, c)
	}

	// 此时栈中元素单调递增，若 k 还没用完的话删掉栈顶元素
	for k > 0 && len(stk) > 0 {
		stk = stk[:len(stk)-1]
		k--
	}
	// 若最后没剩下数字，就是 0
	if len(stk) == 0 {
		return "0"
	}
	// 将栈中字符转化成字符串
	sb := []rune{}
	for len(stk) > 0 {
		sb = append(sb, stk[len(stk)-1])
		stk = stk[:len(stk)-1]
	}
	// 出栈顺序和字符串顺序是反的
	for i, j := 0, len(sb)-1; i < j; i, j = i+1, j-1 {
		sb[i], sb[j] = sb[j], sb[i]
	}
	return string(sb)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRemoveKDigits(t *testing.T) {
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
