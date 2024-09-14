package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
给定一个字符串 s ，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1 。



 示例 1：


输入: s = "leetcode"
输出: 0


 示例 2:


输入: s = "loveleetcode"
输出: 2


 示例 3:


输入: s = "aabb"
输出: -1




 提示:


 1 <= s.length <= 10⁵
 s 只包含小写字母


 Related Topics 队列 哈希表 字符串 计数 👍 754 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func firstUniqChar(s string) int {
	countMap := make(map[int32]int)
	for _, b := range s {
		countMap[b]++
	}
	for i, v := range s {
		if countMap[v] == 1 {
			return i
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFirstUniqueCharacterInAString(t *testing.T) {
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
