package leetcode

import (
	"strings"
	"testing"
	"unicode"
)

/**
如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。

 字母和数字都属于字母数字字符。

 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。



 示例 1：


输入: s = "A man, a plan, a canal: Panama"
输出：true
解释："amanaplanacanalpanama" 是回文串。


 示例 2：


输入：s = "race a car"
输出：false
解释："raceacar" 不是回文串。


 示例 3：


输入：s = " "
输出：true
解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。
由于空字符串正着反着读都一样，所以是回文串。




 提示：


 1 <= s.length <= 2 * 10⁵
 s 仅由可打印的 ASCII 字符组成


 Related Topics 双指针 字符串 👍 754 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
// 将给定的字符串 s 转换成小写字母，并剔除其字母以外的字符
func isPalindrome(s string) bool {
	s = strings.ToLower(s)     // 转换成小写字母
	left, right := 0, len(s)-1 // 一左一右两个指针相向而行
	for left < right {
		if !unicode.IsLetter(rune(s[left])) { // 剔除左指针不是字母的字符
			left++
			continue
		}
		if !unicode.IsLetter(rune(s[right])) { // 剔除右指针不是字母的字符
			right--
			continue
		}
		if s[left] != s[right] { // 比较左右指针是否相等
			return false
		}
		left++
		right--
	}
	return true
}

// 你的代码存在几个小问题：
//
// 1. 只考虑了字母字符，实际要求是剔除所有非字母数字字符，所以应包含数字。
// 2. 字符剔除过程重复两次（一次在 `toLower` 后，一次在循环中），效率不高。
// 3. 强转字符串为 rune，不太必要。
func isPalindrome(s string) bool {
	cleaned := make([]rune, 0, len(s)) // 清理后的字符
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) { // 保留字母或数字
			cleaned = append(cleaned, unicode.ToLower(ch)) // 转换为小写
		}
	}

	left, right := 0, len(cleaned)-1 // 双指针
	for left < right {
		if cleaned[left] != cleaned[right] { // 比较字符是否相等
			return false
		}
		left++
		right--
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidPalindrome(t *testing.T) {

}
