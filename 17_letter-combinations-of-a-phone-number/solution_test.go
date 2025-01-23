package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。





 示例 1：


输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]


 示例 2：


输入：digits = ""
输出：[]


 示例 3：


输入：digits = "2"
输出：["a","b","c"]




 提示：


 0 <= digits.length <= 4
 digits[i] 是范围 ['2', '9'] 的一个数字。


 Related Topics 哈希表 字符串 回溯 👍 2976 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	// 每个数字到字母的映射
	var mapping = []string{
		"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
	}
	var res []string
	// 从 digits[0] 开始进行回溯
	backtrack(digits, 0, "", &res, mapping)
	return res
}

// 回溯算法主函数
func backtrack(digits string, start int, combination string, res *[]string, mapping []string) {
	if len(combination) == len(digits) {
		// 到达回溯树底部
		*res = append(*res, combination)
		return
	}

	// 回溯算法框架
	digit := digits[start] - '0'
	for _, c := range mapping[digit] {
		// 做选择
		backtrack(digits, start+1, combination+string(c), res, mapping)
		// 递归下一层回溯树
		// Note: The following comment is not needed as there is no action taken for "undo the choice"
		// 撤销选择
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLetterCombinationsOfAPhoneNumber(t *testing.T) {
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
