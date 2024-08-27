package leetcode

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

/**
给你一个整数数组 nums 和一个整数 k ，找出 nums 中和至少为 k 的 最短非空子数组 ，并返回该子数组的长度。如果不存在这样的 子数组 ，返回 -1
 。

 子数组 是数组中 连续 的一部分。






 示例 1：


输入：nums = [1], k = 1
输出：1


 示例 2：


输入：nums = [1,2], k = 4
输出：-1


 示例 3：


输入：nums = [2,-1,2], k = 3
输出：3




 提示：


 1 <= nums.length <= 10⁵
 -10⁵ <= nums[i] <= 10⁵
 1 <= k <= 10⁹


 Related Topics 队列 数组 二分查找 前缀和 滑动窗口 单调队列 堆（优先队列） 👍 742 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
type MonotonicQueue struct {
	data []int
}

func NewMonotonicQueue() *MonotonicQueue {
	return &MonotonicQueue{data: []int{}}
}

func (mq *MonotonicQueue) Push(n int) {
	// 保持单调性，移除比新元素大的元素
	for len(mq.data) > 0 && mq.data[len(mq.data)-1] > n {
		mq.data = mq.data[:len(mq.data)-1]
	}
	mq.data = append(mq.data, n)
}

func (mq *MonotonicQueue) Pop() {
	if len(mq.data) > 0 {
		mq.data = mq.data[1:] // 移除队头元素
	}
}

func (mq *MonotonicQueue) Min() int {
	if len(mq.data) > 0 {
		return mq.data[0] // 返回队头元素
	}
	return math.MaxInt64 // 返回无限大以表示无有效的最小值
}

func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	window := NewMonotonicQueue()
	length := math.MaxInt64
	left := 0

	for right := 0; right < len(preSum); right++ {
		window.Push(preSum[right]) // 入队

		// 窗口满足条件，缩小窗口
		for window.Min() != math.MaxInt64 && preSum[right]-window.Min() >= k {
			length = int(math.Min(float64(length), float64(right-left)))
			if window.Min() == preSum[left] {
				window.Pop() // 移除最小值
			}
			left++
		}
	}

	if length == math.MaxInt64 {
		return -1
	}
	return length
}

//leetcode submit region end(Prohibit modification and deletion)

func TestShortestSubarrayWithSumAtLeastK(t *testing.T) {
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
