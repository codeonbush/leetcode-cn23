package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
给你一个整数数组 nums ，和一个表示限制的整数 limit，请你返回最长连续子数组的长度，该子数组中的任意两个元素之间的绝对差必须小于或者等于 limit
。

 如果不存在满足条件的子数组，则返回 0 。



 示例 1：

 输入：nums = [8,2,4,7], limit = 4
输出：2
解释：所有子数组如下：
[8] 最大绝对差 |8-8| = 0 <= 4.
[8,2] 最大绝对差 |8-2| = 6 > 4.
[8,2,4] 最大绝对差 |8-2| = 6 > 4.
[8,2,4,7] 最大绝对差 |8-2| = 6 > 4.
[2] 最大绝对差 |2-2| = 0 <= 4.
[2,4] 最大绝对差 |2-4| = 2 <= 4.
[2,4,7] 最大绝对差 |2-7| = 5 > 4.
[4] 最大绝对差 |4-4| = 0 <= 4.
[4,7] 最大绝对差 |4-7| = 3 <= 4.
[7] 最大绝对差 |7-7| = 0 <= 4.
因此，满足题意的最长子数组的长度为 2 。


 示例 2：

 输入：nums = [10,1,2,4,7,2], limit = 5
输出：4
解释：满足题意的最长子数组是 [2,4,7,2]，其最大绝对差 |2-7| = 5 <= 5 。


 示例 3：

 输入：nums = [4,2,2,2,4,4,2,2], limit = 0
输出：3




 提示：


 1 <= nums.length <= 10^5
 1 <= nums[i] <= 10^9
 0 <= limit <= 10^9


 Related Topics 队列 数组 有序集合 滑动窗口 单调队列 堆（优先队列） 👍 388 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func longestSubarray(nums []int, limit int) int {
	left, right := 0, 0
	windowSize, res := 0, 0
	window := NewMonotonicQueue()

	for right < len(nums) {
		window.Push(nums[right])
		right++
		windowSize++

		for window.Max()-window.Min() > limit {
			window.Pop(nums[left])
			left++
			windowSize--
		}

		res = max(res, windowSize)
	}
	return res
}

func NewMonotonicQueue() *MonotonicQueue {
	return &MonotonicQueue{
		maxQueue: []int{},
		minQueue: []int{},
	}
}

type MonotonicQueue struct {
	maxQueue []int
	minQueue []int
}

func (q *MonotonicQueue) Push(x int) {
	for len(q.maxQueue) > 0 && x > q.maxQueue[len(q.maxQueue)-1] {
		q.maxQueue = q.maxQueue[:len(q.maxQueue)-1]
	}
	q.maxQueue = append(q.maxQueue, x)

	for len(q.minQueue) > 0 && x < q.minQueue[len(q.minQueue)-1] {
		q.minQueue = q.minQueue[:len(q.minQueue)-1]
	}
	q.minQueue = append(q.minQueue, x)
}

func (q *MonotonicQueue) Pop(x int) {
	if len(q.maxQueue) > 0 && x == q.maxQueue[0] {
		q.maxQueue = q.maxQueue[1:]
	}
	if len(q.minQueue) > 0 && x == q.minQueue[0] {
		q.minQueue = q.minQueue[1:]
	}
}

func (q *MonotonicQueue) Max() int {
	return q.maxQueue[0]
}

func (q *MonotonicQueue) Min() int {
	return q.minQueue[0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestContinuousSubarrayWithAbsoluteDiffLessThanOrEqualToLimit(t *testing.T) {
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
