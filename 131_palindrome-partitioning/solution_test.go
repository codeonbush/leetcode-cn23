package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。



 示例 1：


输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]


 示例 2：


输入：s = "a"
输出：[["a"]]




 提示：


 1 <= s.length <= 16
 s 仅由小写英文字母组成


 Related Topics 字符串 动态规划 回溯 👍 1912 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func partition(s string) [][]string {
	var res [][]string
	var track []string
	backtrack(s, 0, &track, &res)
	return res
}

// 回溯算法框架
func backtrack(s string, start int, track *[]string, res *[][]string) {
	if start == len(s) {
		// base case，走到叶子节点
		// 即整个 s 被成功分割为若干个回文子串，记下答案
		tmp := make([]string, len(*track))
		copy(tmp, *track)
		*res = append(*res, tmp)
		return
	}
	for i := start; i < len(s); i++ {
		if !isPalindrome(s, start, i) {
			// s[start..i] 不是回文串，不能分割
			continue
		}
		// s[start..i] 是一个回文串，可以进行分割
		// 做选择，把 s[start..i] 放入路径列表中
		*track = append(*track, s[start:i+1])
		// 进入回溯树的下一层，继续切分 s[i+1..]
		backtrack(s, i+1, track, res)
		// 撤销选择
		*track = (*track)[:len(*track)-1]
	}
}

// 用双指针技巧判断 s[lo..hi] 是否是一个回文串
func isPalindrome(s string, lo, hi int) bool {
	for lo < hi {
		if s[lo] != s[hi] {
			return false
		}
		lo++
		hi--
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPalindromePartitioning(t *testing.T) {
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
