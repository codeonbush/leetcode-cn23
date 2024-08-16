package leetcode

import (
	"testing"
)

/**
给你一个字符串 s，找到 s 中最长的 回文 子串。



 示例 1：


输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。


 示例 2：


输入：s = "cbbd"
输出："bb"




 提示：


 1 <= s.length <= 1000
 s 仅由数字和英文字母组成


 Related Topics 双指针 字符串 动态规划 👍 7305 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		// 以 s[i] 为中心的最长回文子串
		s1 := palindrome(s, i, i)
		// 以 s[i] 和 s[i+1] 为中心的最长回文子串
		s2 := palindrome(s, i, i+1)
		// res = longest(res, s1, s2)
		if len(res) > len(s1) {
			res = res
		} else {
			res = s1
		}
		if len(res) > len(s2) {
			res = res
		} else {
			res = s2
		}
	}
	return res
}

func palindrome(s string, l int, r int) string {
	// 防止索引越界
	for l >= 0 && r < len(s) && s[l] == s[r] {
		// 向两边展开
		l--
		r++
	}
	// 返回以 s[l] 和 s[r] 为中心的最长回文串
	return s[l+1 : r]
}

//runtime:0 ms
//memory:2.5 MB

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestPalindromicSubstring(t *testing.T) {

}
