package leetcode

import (
	"testing"
)

/**
给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几
天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。



 示例 1:


输入: temperatures = [73,74,75,71,69,72,76,73]
输出: [1,1,4,2,1,1,0,0]


 示例 2:


输入: temperatures = [30,40,50,60]
输出: [1,1,1,0]


 示例 3:


输入: temperatures = [30,60,90]
输出: [1,1,0]



 提示：


 1 <= temperatures.length <= 10⁵
 30 <= temperatures[i] <= 100


 Related Topics 栈 数组 单调栈 👍 1807 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	s := make([]int, 0)
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(s) != 0 && temperatures[s[len(s)-1]] <= temperatures[i] {
			s = s[:len(s)-1]
		}
		if len(s) == 0 {
			ans[i] = 0
		} else {
			ans[i] = s[len(s)-1] - i
		}
		s = append(s, i)
	}
	return ans
}

//
//func dailyTemperatures(temperatures []int) []int {
//	n := len(temperatures)
//	ans := make([]int, n)
//	stack := make([]int, 0) // 用栈来存储索引
//
//	for i := 0; i < n; i++ {
//		// 处理当前温度
//		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
//			idx := stack[len(stack)-1]   // 取出栈顶索引
//			stack = stack[:len(stack)-1] // 删除栈顶记录
//			ans[idx] = i - idx           // 更新结果
//		}
//		stack = append(stack, i) // 将当前索引压入栈
//	}
//
//	return ans
//}

//leetcode submit region end(Prohibit modification and deletion)

func TestDailyTemperatures(t *testing.T) {

}
