package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。


 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和
"192.168@1.1" 是 无效 IP 地址。


 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新排序
或删除 s 中的任何数字。你可以按 任何 顺序返回答案。



 示例 1：


输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]


 示例 2：


输入：s = "0000"
输出：["0.0.0.0"]


 示例 3：


输入：s = "101023"
输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]




 提示：


 1 <= s.length <= 20
 s 仅由数字组成


 Related Topics 字符串 回溯 👍 1470 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译。
// 本代码的正确性已通过力扣验证，如有疑问，可以对照 java 代码查看。

import (
	"strconv"
	"strings"
)

// restoreIpAddresses returns all possible valid IP addresses that can be formed by inserting dots into the string s.
func restoreIpAddresses(s string) []string {
	var res []string
	var track []string
	backtrack(s, 0, &track, &res)
	return res
}

// 回溯算法框架
func backtrack(s string, start int, track *[]string, res *[]string) {
	if start == len(s) && len(*track) == 4 {
		// base case，走到叶子节点
		// 即整个 s 被成功分割为合法的四部分，记下答案
		*res = append(*res, strings.Join(*track, "."))
		return
	}
	for i := start; i < len(s); i++ {
		if !isValid(s, start, i) {
			// s[start..i] 不是合法的 ip 数字，不能分割
			continue
		}
		if len(*track) >= 4 {
			// 已经分解成 4 部分了，不能再分解了
			break
		}
		// s[start..i] 是一个合法的 ip 数字，可以进行分割
		// 做选择，把 s[start..i] 放入路径列表中
		*track = append(*track, s[start:i+1])
		// 进入回溯树的下一层，继续切分 s[i+1..]
		backtrack(s, i+1, track, res)
		// 撤销选择
		*track = (*track)[:len(*track)-1]
	}
}

// 判断 s[start..end] 是否是一个合法的 ip 段
func isValid(s string, start, end int) bool {
	length := end - start + 1

	if length == 0 || length > 3 {
		return false
	}

	if length == 1 {
		// 如果只有一位数字，肯定是合法的
		return true
	}

	if s[start] == '0' {
		// 多于一位数字，但开头是 0，肯定不合法
		return false
	}

	if length <= 2 {
		// 排除了开头是 0 的情况，那么如果是两位数，怎么着都是合法的
		return true
	}

	// 现在输入的一定是三位数
	if num, err := strconv.Atoi(s[start : start+length]); err == nil && num > 255 {
		// 不可能大于 255
		return false
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRestoreIpAddresses(t *testing.T) {
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
