package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
给定两个字符串 s 和 t ，它们只包含小写字母。

 字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。

 请找出在 t 中被添加的字母。



 示例 1：


输入：s = "abcd", t = "abcde"
输出："e"
解释：'e' 是那个被添加的字母。


 示例 2：


输入：s = "", t = "y"
输出："y"




 提示：


 0 <= s.length <= 1000
 t.length == s.length + 1
 s 和 t 只包含小写字母


 Related Topics 位运算 哈希表 字符串 排序 👍 481 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func findTheDifference(s string, t string) byte {
	res := 0
	for _, c := range s {
		res ^= int(c)
	}
	for _, d := range t {
		res ^= int(d)
	}

	// 根据异或运算规则，所有字符的异或结果就是多出来的那个字符
	return byte(res)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindTheDifference(t *testing.T) {
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
