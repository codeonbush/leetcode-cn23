package leetcode

import (
	"strings"
	"testing"
)

/**
给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。



 示例 1：


输入：s = "bcabc"
输出："abc"


 示例 2：


输入：s = "cbacdcbc"
输出："acdb"



 提示：


 1 <= s.length <= 10⁴
 s 由小写英文字母组成




 注意：该题与 1081 https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-
characters 相同

 Related Topics 栈 贪心 字符串 单调栈 👍 1094 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicateLetters(s string) string {
	stack := []rune{}
	count := [256]int{}

	// 记录每个字符的数量
	for _, c := range s {
		count[c]++
	}

	inStack := [256]bool{}
	for _, c := range s {
		// 每遍历一个字符，计数减一
		count[c]--

		// 如果字符已存在于堆栈中，跳过
		if inStack[c] {
			continue
		}

		// 按字典序去除更大的元素
		for len(stack) > 0 && stack[len(stack)-1] > c {
			// 若栈顶元素再也不会出现，停止弹出
			if count[stack[len(stack)-1]] == 0 {
				break
			}
			inStack[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}

		// 添加当前字符到堆栈并标记
		stack = append(stack, c)
		inStack[c] = true
	}

	var sb strings.Builder
	// 将结果构建成字符串
	sb.WriteRune(stack[0])
	for i := 1; i < len(stack); i++ {
		sb.WriteRune(stack[i])
	}

	return sb.String()
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRemoveDuplicateLetters(t *testing.T) {

}
