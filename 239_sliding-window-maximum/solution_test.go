package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。


 返回 滑动窗口中的最大值 。



 示例 1：


输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7


 示例 2：


输入：nums = [1], k = 1
输出：[1]




 提示：


 1 <= nums.length <= 10⁵
 -10⁴ <= nums[i] <= 10⁴
 1 <= k <= nums.length


 Related Topics 队列 数组 滑动窗口 单调队列 堆（优先队列） 👍 2856 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
import (
	"container/list"
)

type MonotonicQueue struct {
	maxq list.List
}

func (mq *MonotonicQueue) push(n int) {
	// 将小于 n 的元素全部删除
	for mq.maxq.Len() > 0 && mq.maxq.Back().Value.(int) < n {
		mq.maxq.Remove(mq.maxq.Back())
	}
	// 然后将 n 加入尾部
	mq.maxq.PushBack(n)
}

func (mq *MonotonicQueue) max() int {
	return mq.maxq.Front().Value.(int)
}

func (mq *MonotonicQueue) pop(n int) {
	if n == mq.maxq.Front().Value.(int) {
		mq.maxq.Remove(mq.maxq.Front())
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	window := MonotonicQueue{maxq: list.List{}}
	res := []int{}

	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			//先填满窗口的前 k - 1
			window.push(nums[i])
		} else {
			// 窗口向前滑动，加入新数字
			window.push(nums[i])
			// 记录当前窗口的最大值
			res = append(res, window.max())
			// 移出旧数字
			window.pop(nums[i-k+1])
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSlidingWindowMaximum(t *testing.T) {
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
