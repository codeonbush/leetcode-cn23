package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
给你一份工作时间表 hours，上面记录着某一位员工每天的工作小时数。

 我们认为当员工一天中的工作小时数大于 8 小时的时候，那么这一天就是「劳累的一天」。

 所谓「表现良好的时间段」，意味在这段时间内，「劳累的天数」是严格 大于「不劳累的天数」。

 请你返回「表现良好时间段」的最大长度。



 示例 1：


输入：hours = [9,9,6,0,6,6,9]
输出：3
解释：最长的表现良好时间段是 [9,9,6]。

 示例 2：


输入：hours = [6,6,6]
输出：0




 提示：


 1 <= hours.length <= 10⁴
 0 <= hours[i] <= 16


 Related Topics 栈 数组 哈希表 前缀和 单调栈 👍 540 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func longestWPI(hours []int) int {
	n := len(hours)
	prefixSum := make([]int, n+1)
	stack := []int{0} // 栈初始值为0，表示前缀和初始点
	maxLength := 0

	for i := 1; i <= n; i++ {
		if hours[i-1] > 8 {
			prefixSum[i] = prefixSum[i-1] + 1
		} else {
			prefixSum[i] = prefixSum[i-1] - 1
		}

		// 记录 prefixSum 出现的位置，只保留最早的，栈为单调递减
		if prefixSum[stack[len(stack)-1]] > prefixSum[i] {
			stack = append(stack, i)
		}
	}

	// 从右往左遍历，寻找最大长度的符合条件的子序列
	for i := n; i > 0; i-- {
		for len(stack) > 0 && prefixSum[i] > prefixSum[stack[len(stack)-1]] {
			maxLength = max(maxLength, i-stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestWellPerformingInterval(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{9, 9, 6, 0, 6, 6, 9},
			expected: 3,
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
