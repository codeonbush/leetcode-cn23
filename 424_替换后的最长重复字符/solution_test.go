package leetcode

import (
	"testing"
)

/**
给你一个字符串 s 和一个整数 k 。你可以选择字符串中的任一字符，并将其更改为任何其他大写英文字符。该操作最多可执行 k 次。

 在执行上述操作后，返回 包含相同字母的最长子字符串的长度。



 示例 1：


输入：s = "ABAB", k = 2
输出：4
解释：用两个'A'替换为两个'B',反之亦然。


 示例 2：


输入：s = "AABABBA", k = 1
输出：4
解释：
将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
子串 "BBBB" 有最长重复字母, 答案为 4。
可能存在其他的方法来得到同样的结果。




 提示：


 1 <= s.length <= 10⁵
 s 仅由大写英文字母组成
 0 <= k <= s.length


 Related Topics 哈希表 字符串 滑动窗口 👍 865 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func characterReplacement(s string, k int) int {
	left, right := 0, 0
	// 统计窗口中每个字符的出现次数
	windowCharCount := make([]int, 26)
	// 记录窗口中字符的最多重复次数
	// 记录这个值的意义在于，最划算的替换方法肯定是把其他字符替换成出现次数最多的那个字符
	// 所以窗口大小减去 windowMaxCount 就是所需的替换次数
	windowMaxCount := 0
	// 记录结果长度
	res := 0

	// 开始滑动窗口模板
	for right < len(s) {
		// 扩大窗口
		windowCharCount[s[right]-'A']++
		windowMaxCount = max(windowMaxCount, windowCharCount[s[right]-'A'])
		right++

		for right-left-windowMaxCount > k {
			// 缩小窗口
			windowCharCount[s[left]-'A']--
			left++
			// 这里不用更新 windowMaxCount
			// 因为只有 windowMaxCount 变得更大的时候才可能获得更长的重复子串，才会更新 res
		}
		// 此时一定是一个合法的窗口
		res = max(res, right-left)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestRepeatingCharacterReplacement(t *testing.T) {

}
