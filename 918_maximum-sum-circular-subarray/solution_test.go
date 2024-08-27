package leetcode

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

/**
给定一个长度为 n 的环形整数数组 nums ，返回 nums 的非空 子数组 的最大可能和 。

 环形数组 意味着数组的末端将会与开头相连呈环状。形式上， nums[i] 的下一个元素是 nums[(i + 1) % n] ， nums[i] 的前一个元素
是 nums[(i - 1 + n) % n] 。

 子数组 最多只能包含固定缓冲区 nums 中的每个元素一次。形式上，对于子数组 nums[i], nums[i + 1], ..., nums[j] ，不存在
 i <= k1, k2 <= j 其中 k1 % n == k2 % n 。



 示例 1：


输入：nums = [1,-2,3,-2]
输出：3
解释：从子数组 [3] 得到最大和 3


 示例 2：


输入：nums = [5,-3,5]
输出：10
解释：从子数组 [5,5] 得到最大和 5 + 5 = 10


 示例 3：


输入：nums = [3,-2,2,-3]
输出：3
解释：从子数组 [3] 和 [3,-2,2] 都可以得到最大和 3




 提示：


 n == nums.length
 1 <= n <= 3 * 10⁴
 -3 * 10⁴ <= nums[i] <= 3 * 10⁴


 Related Topics 队列 数组 分治 动态规划 单调队列 👍 723 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
type MonotonicQueue struct {
	q []int
}

func NewMonotonicQueue() *MonotonicQueue {
	return &MonotonicQueue{q: []int{}}
}

// Push adds an element to the queue while maintaining monotonicity
func (mq *MonotonicQueue) Push(n int) {
	// Remove elements from the back until the last element is less than n
	for len(mq.q) > 0 && n < mq.Last() {
		mq.Pop()
	}
	mq.q = append(mq.q, n)
}

// Pop removes the last element
func (mq *MonotonicQueue) Pop() {
	mq.q = mq.q[:len(mq.q)-1]
}

// Min returns the minimum element in the queue
func (mq *MonotonicQueue) Min() int {
	return mq.q[0]
}

// Last returns the last element in the queue
func (mq *MonotonicQueue) Last() int {
	return mq.q[len(mq.q)-1]
}

// Size returns the current size of the queue
func (mq *MonotonicQueue) Size() int {
	return len(mq.q)
}

// Helper function to get the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// maxSubarraySumCircular calculates the maximum sum of any circular subarray
func maxSubarraySumCircular(nums []int) int {
	// Calculate the maximum subarray sum using Kadane's algorithm
	maxKadane := kadane(nums)

	// Calculate the total sum and the minimum subarray sum
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}
	minKadane := kadane(negate(nums))

	// If all numbers are negative, return the maximum found by Kadane's
	if totalSum == -minKadane {
		return maxKadane
	}

	// The maximum circular subarray sum is totalSum - minKadane
	maxCircular := totalSum + minKadane

	// Return the maximum of the two cases
	return max(maxKadane, maxCircular)
}

// kadane implements Kadane's algorithm to find the maximum subarray sum
func kadane(nums []int) int {
	maxSum := math.MinInt32
	currentSum := 0
	for _, num := range nums {
		currentSum = max(num, currentSum+num)
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}

// negate returns a new array with all elements negated
func negate(nums []int) []int {
	negated := make([]int, len(nums))
	for i, num := range nums {
		negated[i] = -num
	}
	return negated
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumSumCircularSubarray(t *testing.T) {
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
