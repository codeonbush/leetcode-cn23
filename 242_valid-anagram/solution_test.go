package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的 字母异位词。



 示例 1:


输入: s = "anagram", t = "nagaram"
输出: true


 示例 2:


输入: s = "rat", t = "car"
输出: false



 提示:


 1 <= s.length, t.length <= 5 * 10⁴
 s 和 t 仅包含小写字母




 进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

 Related Topics 哈希表 字符串 排序 👍 945 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	countMap := make(map[int32]int)
	for _, b := range s {
		countMap[b]++
	}
	countMapt := make(map[int32]int)
	for _, b := range t {
		countMapt[b]++
	}
	for b, count := range countMapt {
		if countMap[b] != count {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidAnagram(t *testing.T) {
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
