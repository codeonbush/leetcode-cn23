package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
有 n 个人排成一个队列，从左到右 编号为 0 到 n - 1 。给你以一个整数数组 heights ，每个整数 互不相同，heights[i] 表示第 i 个
人的高度。

 一个人能 看到 他右边另一个人的条件是这两人之间的所有人都比他们两人 矮 。更正式的，第 i 个人能看到第 j 个人的条件是 i < j 且 min(
heights[i], heights[j]) > max(heights[i+1], heights[i+2], ..., heights[j-1]) 。

 请你返回一个长度为 n 的数组 answer ，其中 answer[i] 是第 i 个人在他右侧队列中能 看到 的 人数 。



 示例 1：




输入：heights = [10,6,8,5,11,9]
输出：[3,1,2,1,1,0]
解释：
第 0 个人能看到编号为 1 ，2 和 4 的人。
第 1 个人能看到编号为 2 的人。
第 2 个人能看到编号为 3 和 4 的人。
第 3 个人能看到编号为 4 的人。
第 4 个人能看到编号为 5 的人。
第 5 个人谁也看不到因为他右边没人。
stk [11 5]
count = 1 len(stk)
count = 1
count = 2
res [1 0]
 示例 2：


输入：heights = [5,1,2,3,10]
输出：[4,1,1,1,0]




 提示：


 n == heights.length
 1 <= n <= 10⁵
 1 <= heights[i] <= 10⁵
 heights 中所有数 互不相同 。


 Related Topics 栈 数组 单调栈 👍 153 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func canSeePersonsCount(heights []int) []int {
	stk := make([]int, 0)
	res := make([]int, len(heights))
	for i := len(heights) - 1; i >= 0; i-- {
		for len(stk) != 0 && stk[len(stk)-1] < heights[i] {

			stk = stk[:len(stk)-1]
			res[i]++
		}
		if len(stk) > 0 {
			res[i]++
		}
		stk = append(stk, heights[i])
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestNumberOfVisiblePeopleInAQueue(t *testing.T) {
	tests := []struct {
		heights  []int
		expected []int
	}{
		{
			heights:  []int{10, 6, 8, 5, 11, 9},
			expected: []int{3, 1, 2, 1, 1, 0},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := canSeePersonsCount(test.heights)
			fmt.Println(result)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For heights %v, expected %v but got %v", test.heights, test.expected, result)
			}
		})
	}
}
