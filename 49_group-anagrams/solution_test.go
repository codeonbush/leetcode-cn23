package leetcode

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

/**
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。



 示例 1:


输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

 示例 2:


输入: strs = [""]
输出: [[""]]


 示例 3:


输入: strs = ["a"]
输出: [["a"]]



 提示：


 1 <= strs.length <= 10⁴
 0 <= strs[i].length <= 100
 strs[i] 仅包含小写字母


 Related Topics 数组 哈希表 字符串 排序 👍 1979 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	// 编码到分组的映射
	codeToGroup := make(map[string][]string)
	for _, s := range strs {
		// 对字符串进行编码
		code := encode(s)
		// 把编码相同的字符串放在一起
		codeToGroup[code] = append(codeToGroup[code], s)
	}

	// 获取结果
	var res [][]string
	for _, group := range codeToGroup {
		res = append(res, group)
	}

	return res
}

// 利用每个字符的出现次数进行编码
func encode(s string) string {
	count := make([]int, 26)
	for _, c := range s {
		count[c-'a']++
	}
	var sb strings.Builder
	for _, num := range count {
		sb.WriteByte(byte(num))
	}
	return sb.String()
}

//leetcode submit region end(Prohibit modification and deletion)

func TestGroupAnagrams(t *testing.T) {
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
