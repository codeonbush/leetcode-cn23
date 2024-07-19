package leetcode

import (
	"testing"
)

/**
给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串， 要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度。

 如果不存在这样的子字符串，则返回 0。



 示例 1：


输入：s = "aaabb", k = 3
输出：3
解释：最长子串为 "aaa" ，其中 'a' 重复了 3 次。


 示例 2：


输入：s = "ababbc", k = 2
输出：5
解释：最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。



 提示：


 1 <= s.length <= 10⁴
 s 仅由小写英文字母组成
 1 <= k <= 10⁵


 Related Topics 哈希表 字符串 分治 滑动窗口 👍 905 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func longestSubstring(s string, k int) int {
	var len int
	for i := 1; i <= 26; i++ {
		// 限制窗口中只能有 i 种不同字符
		len = max(len, longestKLetterSubstr(s, k, i))
	}
	return len
}

// 寻找 s 中含有 count 种字符，且每种字符出现次数都大于 k 的子串
func longestKLetterSubstr(s string, k int, count int) int {
	// 记录答案
	var res int
	// 快慢指针维护滑动窗口，左闭右开区间
	left, right := 0, 0
	// 题目说 s 中只有小写字母，所以用大小 26 的数组记录窗口中字符出现的次数
	windowCount := make([]int, 26)
	// 记录窗口中存在几种不同的字符（字符种类）
	var windowUniqueCount int
	// 记录窗口中有几种字符的出现次数达标（大于等于 k）
	var windowValidCount int
	// 滑动窗口代码模板
	for right < len(s) {
		// 移入字符，扩大窗口
		c := s[right]
		if windowCount[c-'a'] == 0 {
			// 窗口中新增了一种字符
			windowUniqueCount++
		}
		windowCount[c-'a']++
		if windowCount[c-'a'] == k {
			// 窗口中新增了一种达标的字符
			windowValidCount++
		}
		right++

		// 当窗口中字符种类大于 k 时，缩小窗口
		for windowUniqueCount > count {
			// 移出字符，缩小窗口
			d := s[left]
			if windowCount[d-'a'] == k {
				// 窗口中减少了一种达标的字符
				windowValidCount--
			}
			windowCount[d-'a']--
			if windowCount[d-'a'] == 0 {
				// 窗口中减少了一种字符
				windowUniqueCount--
			}
			left++
		}

		// 当窗口中字符种类为 count 且每个字符出现次数都满足 k 时，更新答案
		if windowValidCount == count {
			res = max(res, right-left)
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestSubstringWithAtLeastKRepeatingCharacters(t *testing.T) {

}
