package leetcode

import (
	"testing"
)

/**
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。

 换句话说，s1 的排列之一是 s2 的 子串 。



 示例 1：


输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").


 示例 2：


输入：s1= "ab" s2 = "eidboaoo"
输出：false




 提示：


 1 <= s1.length, s2.length <= 10⁴
 s1 和 s2 仅包含小写字母


 Related Topics 哈希表 双指针 字符串 滑动窗口 👍 1012 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
// 判断 s 中是否存在 t 的排列
func checkInclusion(t string, s string) bool {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for _, c := range []byte(t) {
		need[c]++
	}
	left := 0
	right := 0
	valid := 0
	for right < len(s) {
		c := s[right]
		right++
		// 进行窗口内数据的一系列更新
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		// 判断左侧窗口是否要收缩
		for right-left >= len(t) {
			// 在这里判断是否找到了合法的子串
			if valid == len(need) {
				return true
			}
			d := s[left]
			left++
			// 进行窗口内数据的一系列更新
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	// 未找到符合条件的子串
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPermutationInString(t *testing.T) {

}
