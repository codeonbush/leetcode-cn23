package leetcode

import (
	"testing"
)

/**
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

 异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。



 示例 1:


输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。


 示例 2:


输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。




 提示:


 1 <= s.length, p.length <= 3 * 10⁴
 s 和 p 仅包含小写字母


 Related Topics 哈希表 字符串 滑动窗口 👍 1465 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, t string) []int {
	need := make(map[rune]int)   // 记录目标字符串 t 中每个字符的出现次数
	window := make(map[rune]int) // 记录窗口中每个字符出现的次数
	for _, c := range t {        // 初始化 need
		need[c]++
	}

	left, right := 0, 0
	valid := 0
	res := []int{} // 记录结果
	for right < len(s) {
		c := rune(s[right])
		right++

		if _, ok := need[c]; ok { // 开始滑动窗口，进行窗口数据更新
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for right-left >= len(t) { // 判断左侧窗口是否要收缩
			if valid == len(need) { // 如果当前窗口符合条件，把窗口左侧索引加入 res
				res = append(res, left)
			}
			d := rune(s[left])
			left++

			if _, ok := need[d]; ok { // 进行窗口内数据的更新
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindAllAnagramsInAString(t *testing.T) {

}
